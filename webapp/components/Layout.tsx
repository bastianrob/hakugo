import {FC, PropsWithChildren} from "react";
import Navigation from "./Navbar";

export const Layout: FC<PropsWithChildren> = ({children}) => {
  return (
    <div className="flex flex-col">
      <Navigation />
      <main className="flex grow">{children}</main>
    </div>
  );
};

export default Layout;
