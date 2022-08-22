import "../styles/globals.css";
import CssBaseline from "@mui/material/CssBaseline";

import type {AppProps} from "next/app";
import Layout from "@/components/Layout";
import {
  createTheme,
  StyledEngineProvider,
  ThemeProvider,
} from "@mui/material/styles";

function MyApp({Component, pageProps}: AppProps) {
  const theme = createTheme({
    typography: {
      fontSize: 12,
    },
    components: {
      MuiButton: {
        defaultProps: {
          fullWidth: true,
          variant: "contained",
        },
      },
      MuiTextField: {
        defaultProps: {
          fullWidth: true,
          size: "small",
          InputLabelProps: {
            shrink: true,
          },
        },
      },
    },
  });

  return (
    <StyledEngineProvider injectFirst={false}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </ThemeProvider>
    </StyledEngineProvider>
  );
}

export default MyApp;
