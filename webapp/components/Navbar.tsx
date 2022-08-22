import {
  AppBar,
  Avatar,
  Box,
  Container,
  IconButton,
  Toolbar,
  Tooltip,
} from "@mui/material";
import {FC} from "react";
import CloudCircleIcon from "@mui/icons-material/CloudCircle";
import Spacer from "./Spacer";

export const Navigation: FC = () => {
  return (
    <AppBar position="static">
      <Container maxWidth="xl">
        <Toolbar sx={{gap: 1}}>
          <CloudCircleIcon fontSize="large" />
          <Spacer />

          <Box sx={{flexGrow: 0}}>
            <Tooltip title="Open settings">
              <IconButton size="small">
                <Avatar alt="John Doe" src="https://picsum.photos/40" />
              </IconButton>
            </Tooltip>
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  );
};

export default Navigation;
