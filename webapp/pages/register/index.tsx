import CountrySelect from "@/components/CountrySelect";
import Spacer from "@/components/Spacer";
import {formInputChangeToJson, formSubmitToJson} from "@/libs/formUtilities";
import {
  Box,
  Button,
  Container,
  Divider,
  FormControl,
  FormHelperText,
  Input,
  InputLabel,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import {NextPage} from "next";

const RegisterPage: NextPage = () => {
  const validateForm = (key: string, value: any) => console.log(key, value);
  const submitForm = (data: any) => console.log(data);

  return (
    <Container maxWidth="sm" sx={{py: 2, textAlign: "center", height: "100%"}}>
      <Stack textAlign="center" gap={2}>
        <Typography variant="h5">Sign Up</Typography>

        <form
          onChange={formInputChangeToJson(validateForm)}
          onSubmit={formSubmitToJson(submitForm)}
          className="[&>*]:mb-4"
        >
          <TextField
            id="email"
            name="email"
            label="Email address"
            helperText="We'll never share your email"
            error={false}
            required
            focused
          />

          <TextField
            id="name"
            name="name"
            label="Your name"
            helperText="We'll never share your email"
            required
          />

          <CountrySelect id="country" name="country" required />

          <TextField
            id="phone"
            name="phone"
            label="Your phone number"
            helperText="Please input a valid phone number"
            type="tel"
            required
          />
          <TextField
            id="password"
            name="password"
            type="password"
            label="Password"
            helperText="Password should be minimum of 8 characters"
            required
          />
          <TextField
            type="password"
            id="confirmation"
            name="confirmation"
            label="Password (again)"
            helperText="Confirm your password"
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
