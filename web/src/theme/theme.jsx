import { createTheme } from '@mui/material/styles';

// A custom theme for this app
const mytheme = createTheme({
    palette: {
        mode: 'dark',
        primary: {
            main: '#3F51B5',
        },
        success: {
            main: '#388e3c',
        },
    },
});

export default mytheme;
