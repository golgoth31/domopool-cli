import * as React from "react";
import dataprovider from '../dataprovider/dataprovider';
import {
  Formik,
  Form,
  Field,
} from 'formik';
import {
  TextField,
  Checkbox,
} from 'formik-material-ui';
import SubmitButton from "./SubmitButton";
import CircularProgress from '@material-ui/core/CircularProgress';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import isIp from 'is-ip';
import NetworkFields from './NetworkForm';
import GlobalFields from './GlobalForm';

// export const ConfigForm: React.SFC = () => {
export default class ConfigForm extends React.Component {
  constructor(
    private conf: any
  ) { super(conf); };

  state = {
    config: this.conf,
    ipDisabled: false,
  }

  componentDidMount() {
    dataprovider.get(`config`)
      .then(res => {
        const resp = res.data;
        this.setState({ config: resp });
      })
  }

  // private validateEmail(value) {
  //   let error;
  //   if (!value) {
  //     error = 'Required';
  //   } else if (!/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(value)) {
  //     error = 'Invalid email address';
  //   }
  //   return error;
  // }


  private validateIP = (value: any, state: any) => {
    let error;
    if (!this.state.ipDisabled) {
      if (!value) {
        error = 'Required';
      } else if (!isIp(value)) {
        error = 'Invalid IP format';
      }
    }
    return error;
  }
  private handleDHCPClick = (event: React.ChangeEvent<HTMLInputElement>) => {
    this.state.ipDisabled = event.target.checked
  }

  render() {
    if (this.state.config.global != undefined) {
      this.state.ipDisabled = this.state.config.network.dhcp;
      return (
        <Formik
          enableReinitialize
          initialValues={{
            global: {
              lcdbacklightduration: this.state.config.global.lcdbacklightduration,
            },
            network: {
              dhcp: this.state.config.network.dhcp,
              ip: this.state.config.network.ip,
              netmask: this.state.config.network.netmask,
              gateway: this.state.config.network.gateway,
              dns: this.state.config.network.dns,
            }
          }}
          onSubmit={(values, actions) => {
            setTimeout(() => {
              alert(JSON.stringify(values, null, 2));
              actions.setSubmitting(false);
            }, 1000);
          }}
        >
          {props => (

            <Form>
              <GlobalFields {...props} />
              <br />

              <NetworkFields {...props} />
              <br />
              <SubmitButton />
            </Form>

          )}
        </Formik >
      )
    } else {
      return (
        <CircularProgress />
      )
    }
  };
};
