import * as React from "react";
import * as ReactDOM from "react-dom";
import { ThemeProvider } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Dashboard from './components/Dashboard';
import theme from './theme';


ReactDOM.render(
  <ThemeProvider theme={theme}>
    < CssBaseline />
    <Dashboard />
  </ThemeProvider>,
  document.querySelector('#root'),
);
