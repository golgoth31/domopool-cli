// import * as React from "react";
import * as ReactDOM from "react-dom";
import { BrowserRouter } from 'react-router-dom';
import { ThemeProvider } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Dashboard from './layouts/Dashboard';
import theme from './theme/theme';

// declare global {
//     interface Window {
//         DOMOPOOL_HOST: string;
//         DOMOPOOL_PORT: string;
//         DOMOPOOL_SCHEME: string;
//     }

ReactDOM.render(
    <BrowserRouter>
        <ThemeProvider theme={theme}>
            < CssBaseline />
            <Dashboard />
        </ThemeProvider>
    </BrowserRouter>,
    document.querySelector('#root'),
);
