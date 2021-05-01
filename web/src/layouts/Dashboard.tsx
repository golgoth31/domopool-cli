import * as React from "react";
import clsx from 'clsx';
import Drawer from '@material-ui/core/Drawer';
import Divider from '@material-ui/core/Divider';
import IconButton from '@material-ui/core/IconButton';
import Container from '@material-ui/core/Container';
import List from '@material-ui/core/List';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import MainListItems from './Menu';
import TopAppBar from './AppBar';
import Routes from '../routes';
import useStyles from '../theme/useStyles';

// export default function Dashboard() {
export default function Dashboard() {
    const [open, setOpen] = React.useState(true);
    const classes = useStyles();
    const handleDrawerOpen = () => {
        setOpen(true);
    };
    const handleDrawerClose = () => {
        setOpen(false);
    };

    return (
        <div className={classes.root} >
            <TopAppBar handleDrawerOpen={handleDrawerOpen} open={open} />
            <Drawer
                variant="permanent"
                classes={{
                    paper: clsx(classes.drawerPaper, !open && classes.drawerPaperClose),
                }}
                open={open}
            >
                <div className={classes.toolbarIcon}>
                    <IconButton onClick={handleDrawerClose}>
                        <ChevronLeftIcon />
                    </IconButton>
                </div>
                <Divider />
                <List>
                    <MainListItems handleDrawerClose={handleDrawerClose} />
                </List>
            </Drawer>
            <main className={classes.content}>
                <div className={classes.appBarSpacer} />
                <Container maxWidth="lg" className={classes.container}>
                    <Routes />
                </Container>
            </main>
        </div>
    );
}
