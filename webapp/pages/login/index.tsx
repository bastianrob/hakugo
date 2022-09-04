import {Button, Container, Stack, TextField, Typography} from "@mui/material";
import {NextPage} from "next";
import {useRef, useState} from "react";
import {
  formInputChangeToJson,
  formSubmitToJson,
  formToJson,
} from "@/libs/formUtilities";
import {extendedJoi, validationErrorToJson} from "@/libs/extendedJoi";
import {useRouter} from "next/router";
import {useCustomerLoginMutation} from "graphql/generated";
import {setCookie} from "cookies-next";
import {useSnackbar} from "notistack";

interface errorHelper {
  email: string;
  password: string;
}

const errorHelperInitialValue: errorHelper = {
  email: "",
  password: "",
};

const LoginPage: NextPage = () => {
  const router = useRouter();
  const formRef = useRef<HTMLFormElement>(null);
  const {enqueueSnackbar, closeSnackbar} = useSnackbar();
  const [errorHelper, setErrorHelper] = useState<errorHelper>(
    errorHelperInitialValue,
  );

  const [doLogin, {loading}] = useCustomerLoginMutation();

  const validateForm = (data: any): undefined | any => {
    const schema = extendedJoi.object({
      email: extendedJoi
        .string()
        .email({tlds: {allow: false}})
        .message("Please input a valid email address")
        .required(),
      password: extendedJoi
        .string()
        .min(8)
        .message("Password should be at least 8 characters long")
        .required(),
    });

    const {error, value} = schema.validate(data, {abortEarly: false});
    const errorJson = validationErrorToJson<errorHelper>(
      errorHelperInitialValue,
      error,
    );

    if (error) {
      setErrorHelper(errorJson);
      return;
    }

    setErrorHelper(errorHelperInitialValue);
    return value;
  };

  const submitForm = async (data: any) => {
    const validatedFormValue = validateForm(data);
    if (!validatedFormValue) return;

    const {email, password} = validatedFormValue;

    try {
      const {data} = await doLogin({
        variables: {
          user: email as string,
          pass: password as string,
        },
      });
      setCookie("access-token", data?.customerLogin?.accessToken);
      setCookie("logged-user", `${email}`);
      const snackId = enqueueSnackbar("You have successfully logged in", {
        variant: "success",
        onClick: () => closeSnackbar(snackId),
      });

      setTimeout(() => router.replace("/"), 3000);
    } catch (error: any) {
      const snackId = enqueueSnackbar(error.message, {
        variant: "error",
        onClick: () => closeSnackbar(snackId),
      });
    }
  };

  return (
    <Container
      maxWidth="sm"
      className="shadow py-4 text-center h-full rounded mt-4"
    >
      <Stack dir="vertical" gap={4}>
        <Typography variant="h5">Sign In</Typography>

        <form
          ref={formRef}
          onChange={formInputChangeToJson(validateForm)}
          onSubmit={formSubmitToJson(submitForm)}
          className="[&>*]:mb-4"
        >
          <TextField
            id="email"
            name="email"
            label="Email address"
            helperText={
              errorHelper.email || "Please input a valid email address"
            }
            error={!!errorHelper.email}
            required
            focused
          />

          <TextField
            id="password"
            name="password"
            type="password"
            label="Password"
            helperText={
              errorHelper.password ||
              "Password should be minimum of 8 characters"
            }
            error={!!errorHelper.password}
            required
          />
          <Button disabled={loading} type="submit">
            Sign In
          </Button>
        </form>
      </Stack>
    </Container>
  );
};

export default LoginPage;
