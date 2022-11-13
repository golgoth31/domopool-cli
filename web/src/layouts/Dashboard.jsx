import * as React from "react";
import {
    Container
} from '@mui/material';
import { Outlet } from "react-router-dom";
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import CssBaseline from '@mui/material/CssBaseline';
import TopAppBar from './AppBar';
import AppRoutes from '../routes';
// import useStyles from '../theme/useStyles.jsx.old';
import SideDrawer from "./Drawer";

// export default function Dashboard() {
//     const [open, setOpen] = React.useState(true);
//     const classes = useStyles();
//     const toggleDrawer = () => {
//         setOpen(!open);
//     };

//     return (
//         <div className={classes.root} >
//             <TopAppBar toggleDrawer={toggleDrawer} open={open} />
//             <SideDrawer toggleDrawer={toggleDrawer} open={open} />

//             <main className={classes.content}>
//                 <div className={classes.appBarSpacer} />
//                 <Container maxWidth="lg" className={classes.container}>
//                     <Routes />
//                 </Container>
//             </main>
//         </div>
//     );
// }

function DashboardContent() {
    const [open, setOpen] = React.useState(true);
    const toggleDrawer = () => {
        setOpen(!open);
    };

    return (
        <Box sx={{ display: 'flex' }}>
            <CssBaseline />

            <TopAppBar toggleDrawer={toggleDrawer} open={open} />
            <SideDrawer toggleDrawer={toggleDrawer} open={open} />

            <Box
                component="main"
                sx={{
                    backgroundColor: (theme) =>
                        theme.palette.mode === 'light'
                            ? theme.palette.grey[100]
                            : theme.palette.grey[900],
                    flexGrow: 1,
                    height: '100vh',
                    overflow: 'auto',
                }}
            >
                <Toolbar />
                <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
                    <AppRoutes />
                    <Outlet />
                </Container>
            </Box>
        </Box>
    );
}

export default function Dashboard() {
    return <DashboardContent />;
}
