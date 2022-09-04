import {
  AppBar,
  Avatar,
  Box,
  Button,
  Container,
  Divider,
  IconButton,
  ListItemIcon,
  Menu,
  MenuItem,
  Toolbar,
  Tooltip,
} from "@mui/material";
import {FC, useRef, useState} from "react";
import CloudCircleIcon from "@mui/icons-material/CloudCircle";
import Spacer from "./Spacer";
import {useGetCurrentUserLazyQuery} from "graphql/generated";
import {useEffectOnce, useToggle} from "react-use";
import {getInitialFromEmail} from "@/libs/stringUtilities";
import {Logout, Person, PersonAdd, Settings} from "@mui/icons-material";
import clsx from "clsx";

export const Navigation: FC = () => {
  const [isLoggedIn, setLoggedIn] = useState(false);
  const [initial, setInitial] = useState("");
  const [isOpen, setOpened] = useToggle(false);
  const menuAnchor = useRef(null);

  const [getCurrentUser] = useGetCurrentUserLazyQuery();

  useEffectOnce(() => {
    getCurrentUser().then(({data}) => {
      if (data?.currentUser?.isLoggedIn) {
        setLoggedIn(true);
        setInitial(getInitialFromEmail(data.currentUser.email!));
      }
    });
  });

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
                onClick={() => setOpened(true)}
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
              <Button variant="text" color="inherit">
                Sign In
              </Button>
            )}
          </Box>

          <Menu
            open={isOpen}
            anchorEl={menuAnchor?.current}
            onClose={() => setOpened(false)}
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
            <MenuItem>
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
