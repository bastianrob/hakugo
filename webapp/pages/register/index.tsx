import CountrySelect, {CountryType} from "@/components/CountrySelect";
import Spacer from "@/components/Spacer";
import {extendedJoi, validationErrorToJson} from "@/libs/extendedJoi";
import {formInputChangeToJson, formSubmitToJson} from "@/libs/formUtilities";
import {
  Button,
  Container,
  InputAdornment,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import {ValidationError} from "joi";
import {NextPage} from "next";
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
  const [countryType, setCountryType] = useState<CountryType>(emptyCountry);
  const [errorHelper, setErrorHelper] = useState<errorHelper>(
    errorHelperInitialValue,
  );

  const validateInputChange = (key: string, value: any) => {
    if (!key) return;

    let validationError: ValidationError | undefined = undefined;
    switch (key) {
      case "name": {
        const {error} = extendedJoi
          .string()
          .min(2)
          .message("Name should be longer than 2 characters")
          .required()
          .validate(value);
        validationError = error;
        break;
      }
      case "email": {
        const {error} = extendedJoi
          .string()
          .email({tlds: {allow: false}})
          .message("Please input a valid email address")
          .required()
          .validate(value);
        validationError = error;
        break;
      }
      case "phone": {
        const {error} = extendedJoi
          .string()
          .phoneNumber({
            defaultCountry: countryType.code || undefined,
            format: "e164",
            strict: countryType.code ? true : false,
          })
          .message("Please input a valid phone number")
          .validate(value);
        validationError = error;
        break;
      }
      case "password": {
        const {error} = extendedJoi
          .string()
          .min(8)
          .message("Password should be at least 8 characters long")
          .required()
          .validate(value);
        validationError = error;
        break;
      }
    }

    if (!validationError) {
      setErrorHelper({...errorHelper, [key]: ""});
    } else {
      const errorJson = validationErrorToJson(
        errorHelper,
        validationError,
        key,
      );
      setErrorHelper(errorJson);
    }
  };

  const submitForm = (data: any) => {
    if (!countryType.code) {
      setErrorHelper({
        ...errorHelperInitialValue,
        country: "Please input your phone number's country of origin",
      });
      return;
    }

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
      phone: extendedJoi
        .string()
        .phoneNumber({
          defaultCountry: countryType.code || undefined,
          format: "e164",
          strict: countryType.code ? true : false,
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
    } else {
      setErrorHelper(errorHelperInitialValue);
    }

    // console.debug(value);
  };

  const handleCountryChanged = (_: any, value: any) => {
    if (value) setCountryType(value as CountryType);
    else setCountryType(emptyCountry);
  };

  return (
    <Container maxWidth="sm" sx={{py: 2, textAlign: "center", height: "100%"}}>
      <Stack textAlign="center" gap={2}>
        <Typography variant="h5">Sign Up</Typography>

        <form
          onChange={formInputChangeToJson(validateInputChange)}
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

          <Button type="submit">Sign Up</Button>
        </form>
      </Stack>
    </Container>
  );
};

export default RegisterPage;
