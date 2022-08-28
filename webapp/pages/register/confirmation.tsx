import apolloClient, {GraphqlErrorCodes} from "@/libs/apolloClient";
import {GraphQLErrors, NetworkError} from "@apollo/client/errors";
import {
  Alert,
  AlertColor,
  AlertTitle,
  Button,
  Container,
  Stack,
  Typography,
} from "@mui/material";
import {
  AuthenticationVerifyDocument,
  AuthenticationVerifyMutation,
} from "graphql/generated";
import {GetServerSidePropsContext, NextPage} from "next";
import {useRouter} from "next/router";
import {useMemo} from "react";
import {useTimeoutFn} from "react-use";

type VerificationState =
  | "waiting_verification"
  | "already_verified"
  | "already_expired"
  | "unexpected_error";

interface ConfirmationPageProps {
  email: string;
  verificationState: VerificationState;
}

const ConfirmationPage: NextPage<ConfirmationPageProps> = ({
  email,
  verificationState,
}) => {
  const router = useRouter();
  const [severity, title, desc] = useMemo<[AlertColor, string, string]>(() => {
    switch (verificationState) {
      case "waiting_verification":
        return [
          "info",
          `We have sent you an email to '${email}'`,
          "Please check your inbox or spam folder and click on the confirmation link we've sent to you",
        ];
      case "already_verified":
        return [
          "success",
          `your email: '${email}' is already verified`,
          "We will redirect you to homepage in 3 seconds...",
        ];
      case "already_expired":
        return [
          "warning",
          `Your verification code is already expired`,
          "Please click on [Resend Verification] button below to rety the verification process",
        ];
      case "unexpected_error":
        return [
          "error",
          "Unknown error ocurred",
          "Please refresh this page to retry",
        ];
    }
  }, [verificationState]);

  useTimeoutFn(() => {
    if (verificationState !== "already_verified") return;
    router.replace("/");
  }, 3 * 1000);

  return (
    <Container maxWidth="sm">
      <Stack dir="vertical" gap={4}>
        <Typography variant="h6"></Typography>
        <Alert severity={severity}>
          <AlertTitle className="font-bold">{title}</AlertTitle>
          {desc}
        </Alert>

        <Button
          disabled={
            verificationState === "already_verified" ||
            verificationState === "unexpected_error"
          }
        >
          Resend Verification
        </Button>
      </Stack>
    </Container>
  );
};

export default ConfirmationPage;

export const getServerSideProps = async (
  context: GetServerSidePropsContext,
) => {
  let verificationState: VerificationState = "waiting_verification";
  const {email, code} = context.query;

  if (email && code) {
    try {
      const {data} = await apolloClient.mutate<AuthenticationVerifyMutation>({
        mutation: AuthenticationVerifyDocument,
        variables: {email, code},
      });

      // TODO: get accessToken and automatically login
      console.log(data?.authenticationVerify?.accessToken);
      verificationState = "already_verified";
    } catch (error: any) {
      const {graphQLErrors, networkError} = error as {
        graphQLErrors?: GraphQLErrors;
        networkError?: NetworkError;
      };

      if (graphQLErrors && graphQLErrors.length > 0) {
        const [
          {
            extensions: {code},
          },
        ] = graphQLErrors;

        switch (code) {
          case GraphqlErrorCodes.NotFound:
          case GraphqlErrorCodes.InputMalformed:
          case GraphqlErrorCodes.UnexpectedError:
            verificationState = "unexpected_error";
            break;
          case "EXPIRED":
            verificationState = "already_expired";
            break;
          case "USED":
            verificationState = "already_verified";
            break;
        }
      } else if (networkError) {
        verificationState = "unexpected_error";
      }
    }
  }

  return {
    props: {
      email: (email as string) || "",
      verificationState,
    },
  };
};
