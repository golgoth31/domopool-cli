// import * as React from "react";
import * as ReactDOM from "react-dom";
import { BrowserRouter } from 'react-router-dom';
import { ThemeProvider } from '@mui/material/styles';
import Box from '@mui/material/Box';
import CssBaseline from '@mui/material/CssBaseline';
import Dashboard from './layouts/Dashboard';
import mytheme from './theme/theme';

// declare global {
//     interface Window {
//         DOMOPOOL_HOST: string;
//         DOMOPOOL_PORT: string;
//         DOMOPOOL_SCHEME: string;
//     }

ReactDOM.render(
    <BrowserRouter>
        <ThemeProvider theme={mytheme}>
            {/* <Box sx={{ display: 'flex' }}>
                <CssBaseline /> */}
            <Dashboard />
            {/* </Box> */}
        </ThemeProvider>
    </BrowserRouter>,
    document.querySelector('#root'),
);
