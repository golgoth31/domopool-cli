import * as React from "react";
import * as ReactDOM from "react-dom";
import { ThemeProvider } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Dashboard from './components/Dashboard';
import theme from './theme';

declare global {
  interface Window {
    DOMOPOOL_HOST: string;
    DOMOPOOL_PORT: string;
    DOMOPOOL_SCHEME: string;
  }
}

ReactDOM.render(
  <ThemeProvider theme={theme}>
    < CssBaseline />
    <Dashboard />
  </ThemeProvider>,
  document.querySelector('#root'),
);
