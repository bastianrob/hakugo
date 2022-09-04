import {
  AppBar,
  Avatar,
  Box,
  Button,
  Container,
  IconButton,
  ListItemIcon,
  Menu,
  MenuItem,
  Toolbar,
} from "@mui/material";
import {FC, useRef, useState} from "react";
import CloudCircleIcon from "@mui/icons-material/CloudCircle";
import Spacer from "./Spacer";
import {useGetCurrentUserLazyQuery} from "graphql/generated";
import {useEffectOnce, useToggle} from "react-use";
import {getInitialFromEmail} from "@/libs/stringUtilities";
import {Logout, Person, Settings} from "@mui/icons-material";
import clsx from "clsx";
import {useRouter} from "next/router";
import {deleteCookie} from "cookies-next";

export const Navigation: FC = () => {
  const [initial, setInitial] = useState("");
  const [isLoggedIn, setLoggedIn] = useState(false);
  const [isMenuOpen, setMenuOpened] = useToggle(false);
  const [getCurrentUser] = useGetCurrentUserLazyQuery();

  const router = useRouter();
  const menuAnchor = useRef(null);

  useEffectOnce(() => {
    getCurrentUser().then(({data}) => {
      if (data?.currentUser?.isLoggedIn) {
        setLoggedIn(true);
        setInitial(getInitialFromEmail(data.currentUser.email!));
      }
    });
  });

  const handleLogout = () => {
    deleteCookie("access-token");
    deleteCookie("logged-user");
    setMenuOpened(false);
    setLoggedIn(false);
    setInitial("");

    router.replace("/");
  };

  return (
    <AppBar position="static">
      <Container maxWidth="lg">
        <Toolbar sx={{gap: 1}}>
          <CloudCircleIcon fontSize="large" />
          <Spacer />

          <Box sx={{flexGrow: 0}}>
            {isLoggedIn ? (
              <IconButton
                ref={menuAnchor}
                size="small"
                onClick={() => setMenuOpened(true)}
              >
                <Avatar
                  className={clsx(
                    "font-bold text-sm h-[32px] w-[32px]",
                    isLoggedIn && "bg-orange-400",
                  )}
                >
                  {initial}
                </Avatar>
              </IconButton>
            ) : (
              <Button
                variant="text"
                color="inherit"
                onClick={() => router.push("/login")}
              >
                Sign In
              </Button>
            )}
          </Box>

          <Menu
            open={isMenuOpen}
            anchorEl={menuAnchor?.current}
            onClose={() => setMenuOpened(false)}
            anchorOrigin={{
              vertical: "bottom",
              horizontal: "right",
            }}
            transformOrigin={{
              vertical: "top",
              horizontal: "right",
            }}
          >
            <MenuItem>
              <ListItemIcon>
                <Person fontSize="small" />
              </ListItemIcon>
              My Account
            </MenuItem>
            <MenuItem>
              <ListItemIcon>
                <Settings fontSize="small" />
              </ListItemIcon>
              Settings
            </MenuItem>
            <MenuItem onClick={handleLogout}>
              <ListItemIcon>
                <Logout fontSize="small" />
              </ListItemIcon>
              Logout
            </MenuItem>
          </Menu>
        </Toolbar>
      </Container>
    </AppBar>
  );
};

export default Navigation;
