import {SnackbarProvider} from "notistack";
import {FC, PropsWithChildren} from "react";
import Navigation from "./Navbar";

export const Layout: FC<PropsWithChildren> = ({children}) => {
  return (
    <div className="flex flex-col">
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
        <main className="flex grow">{children}</main>
      </SnackbarProvider>
    </div>
  );
};

export default Layout;
