import CountrySelect, {CountryType} from "@/components/CountrySelect";
import Spacer from "@/components/Spacer";
import {extendedJoi, validationErrorToJson} from "@/libs/extendedJoi";
import {
  formInputChangeToJson,
  formSubmitToJson,
  formToJson,
} from "@/libs/formUtilities";
import {
  Button,
  Container,
  Input,
  InputAdornment,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import {useCustomerRegisterMutation} from "graphql/generated";
import {NextPage} from "next";
import {useRouter} from "next/router";
import {useSnackbar} from "notistack";
import {useRef} from "react";
import {useState} from "react";

const emptyCountry: CountryType = {label: "", code: "", phone: "0"};

interface errorHelper {
  name: string;
  email: string;
  country: string;
  phone: string;
  password: string;
  confirmation: string;
}

const errorHelperInitialValue: errorHelper = {
  name: "",
  email: "",
  country: "",
  phone: "",
  password: "",
  confirmation: "",
};

const RegisterPage: NextPage = () => {
  const router = useRouter();
  const {enqueueSnackbar, closeSnackbar} = useSnackbar();

  const formRef = useRef<HTMLFormElement>(null);
  const [countryType, setCountryType] = useState<CountryType>(emptyCountry);
  const [errorHelper, setErrorHelper] = useState<errorHelper>(
    errorHelperInitialValue,
  );

  const [register] = useCustomerRegisterMutation();

  const validateForm = (data: any): undefined | any => {
    const schema = extendedJoi.object({
      email: extendedJoi
        .string()
        .email({tlds: {allow: false}})
        .message("Please input a valid email address")
        .required(),
      name: extendedJoi
        .string()
        .min(2)
        .message("Name should be longer than 2 characters")
        .required(),
      country: extendedJoi.string().required().messages({
        "string.empty": "Please input your phone number's country of origin",
      }),
      phone: extendedJoi
        .string()
        .phoneNumber({
          format: "e164",
          defaultCountry: data.country || undefined,
          strict: true,
        })
        .message("Please input a valid phone number"),
      password: extendedJoi
        .string()
        .min(8)
        .message("Password should be at least 8 characters long")
        .required(),
      confirmation: extendedJoi
        .any()
        .valid(extendedJoi.ref("password"))
        .required()
        .messages({"any.only": "Confirmation password does not match"}),
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

    const {country, ...registrationInput} = validatedFormValue;

    try {
      await register({
        variables: {
          input: {
            ...registrationInput,
            provider: "email",
          },
        },
      });
    } catch (error) {
      const snackId = enqueueSnackbar((error as Error).message, {
        variant: "error",
        onClick: () => closeSnackbar(snackId),
      });
      return;
    }

    router.replace({
      pathname: "/register/confirmation",
      query: {
        email: validatedFormValue.email,
      },
    });
  };

  const handleCountryChanged = (_: any, value: any) => {
    const data = formToJson(formRef.current!);
    if (value) {
      setCountryType(value);
      validateForm({
        ...data,
        country: value.code,
      });
    } else {
      setCountryType(emptyCountry);
      validateForm({
        ...data,
        country: "",
      });
    }
  };

  const canBeSubmitted = () => {
    return (
      !errorHelper.name &&
      !errorHelper.email &&
      !errorHelper.country &&
      !errorHelper.phone &&
      !errorHelper.password &&
      !errorHelper.confirmation
    );
  };

  return (
    <Container maxWidth="sm" className="shadow py-4 text-center h-full rounded">
      <Stack textAlign="center" gap={2}>
        <Typography variant="h5">Sign Up</Typography>

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
            id="name"
            name="name"
            label="Your name"
            helperText={errorHelper.name || "Please tell us your name"}
            error={!!errorHelper.name}
            required
          />

          <CountrySelect
            onChange={handleCountryChanged}
            helperText={
              errorHelper.country || "Choose your phone number's country"
            }
            error={!!errorHelper.country}
            required
          />

          <Input name="country" type="hidden" value={countryType.code} />

          <TextField
            id="phone"
            name="phone"
            label="Your phone number"
            helperText={
              errorHelper.phone || "Please input a valid phone number"
            }
            error={!!errorHelper.phone}
            type="tel"
            required
            InputProps={{
              sx: {alignItems: "baseline"},
              startAdornment: (
                <InputAdornment position="start">
                  +{countryType.phone}
                </InputAdornment>
              ),
            }}
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
          <TextField
            type="password"
            id="confirmation"
            name="confirmation"
            label="Password (again)"
            helperText={errorHelper.confirmation || "Confirm your password"}
            error={!!errorHelper.confirmation}
            required
          />

          <Spacer />

          <Button disabled={!canBeSubmitted()} type="submit">
            Sign Up
          </Button>
        </form>
      </Stack>
    </Container>
  );
};

export default RegisterPage;
