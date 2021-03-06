import { Link } from 'react-router-dom';
import { List, ListItem, ListItemIcon, ListItemText } from '@material-ui/core/';
import DashboardIcon from '@material-ui/icons/Dashboard';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';
import PeopleIcon from '@material-ui/icons/People';
import HomeIcon from '@material-ui/icons/Home';

const MenuLayout = (props: any) => {
    return (
        <List>
            <ListItem button component={Link} to="/" onClick={props.handleDrawerClose}>
                <ListItemIcon>
                    <HomeIcon />
                </ListItemIcon>
                <ListItemText primary="Home" />
            </ListItem>

            <ListItem button component={Link} to="/config" onClick={props.handleDrawerClose}>
                <ListItemIcon>
                    <DashboardIcon />
                </ListItemIcon>
                <ListItemText primary="Config" />
            </ListItem>

            <ListItem button>
                <ListItemIcon>
                    <ShoppingCartIcon />
                </ListItemIcon>
                <ListItemText primary="Data" />
            </ListItem>

            <ListItem button component={Link} to="/alarms" onClick={props.handleDrawerClose}>
                <ListItemIcon>
                    <PeopleIcon />
                </ListItemIcon>
                <ListItemText primary="Alarms" />
            </ListItem>
        </List>
    )
};

export default MenuLayout;
