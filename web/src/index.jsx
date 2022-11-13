// import * as React from "react";
// import * as ReactDOM from "react-dom";
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import { ThemeProvider } from '@mui/material/styles';
import Dashboard from './layouts/Dashboard';
import mytheme from './theme/theme';

// declare global {
//     interface Window {
//         DOMOPOOL_HOST: string;
//         DOMOPOOL_PORT: string;
//         DOMOPOOL_SCHEME: string;
//     }

// ReactDOM.render(
//     <BrowserRouter>
//         <ThemeProvider theme={mytheme}>
//             {/* <Box sx={{ display: 'flex' }}>
//                 <CssBaseline /> */}
//             <Dashboard />
//             {/* </Box> */}
//         </ThemeProvider>
//     </BrowserRouter>,
//     document.querySelector('#root'),
// );

// import { createRoot } from 'react-dom/client';
const container = document.getElementById('root');
const root = createRoot(container); // createRoot(container!) if you use TypeScript
root.render(
    <BrowserRouter>
        <ThemeProvider theme={mytheme}>
            {/* <Box sx={{ display: 'flex' }}>
                <CssBaseline /> */}
            <Dashboard />
            {/* </Box> */}
        </ThemeProvider>
    </BrowserRouter>
);
