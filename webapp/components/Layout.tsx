import {SnackbarProvider} from "notistack";
import {FC, PropsWithChildren} from "react";
import Navigation from "./Navbar";

export const Layout: FC<PropsWithChildren> = ({children}) => {
  return (
    <>
      <Navigation />
      <SnackbarProvider
        dense
        preventDuplicate
        autoHideDuration={8000}
        anchorOrigin={{
          vertical: "bottom",
          horizontal: "center",
        }}
      >
        <main>{children}</main>
      </SnackbarProvider>
    </>
  );
};

export default Layout;
