import {
  AppBar,
  Avatar,
  Box,
  Container,
  IconButton,
  Toolbar,
  Tooltip,
} from "@mui/material";
import {FC, useState} from "react";
import CloudCircleIcon from "@mui/icons-material/CloudCircle";
import Spacer from "./Spacer";
import {useGetCurrentUserLazyQuery} from "graphql/generated";
import {useEffectOnce} from "react-use";
import {getInitialFromEmail} from "@/libs/stringUtilities";

export const Navigation: FC = () => {
  const [initial, setInitial] = useState("");
  const [getCurrentUser] = useGetCurrentUserLazyQuery();

  useEffectOnce(() => {
    getCurrentUser().then(({data}) => {
      if (data?.currentUser?.isLoggedIn) {
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
            <IconButton size="small">
              <Avatar className="bg-orange-500 font-bold text-sm h-[32px] w-[32px]">
                {initial}
              </Avatar>
            </IconButton>
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  );
};

export default Navigation;
