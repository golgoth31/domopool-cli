import * as React from "react";
import dataprovider from '../dataprovider/dataprovider';
import {
  Formik,
  Form,
} from 'formik';
import SubmitButton from "./SubmitButton";
import CircularProgress from '@material-ui/core/CircularProgress';
import NetworkFields from './NetworkForm';
import GlobalFields from './GlobalForm';
import TwoutFields from './TwoutForm';
import TwinFields from './TwinForm';
import TambFields from './TambForm';
import { Grid } from "@material-ui/core";

var domopool_pb = require('./../proto/domopool_pb');
// import { } from './../proto/domopool_pb';

export default class ConfigForm extends React.Component {
  constructor(
    private conf: any
  ) { super(conf); };

  state = {
    config: this.conf,
    ipDisabled: false,
  }

  componentDidMount() {
    dataprovider.get(`/api/v1/config`, {
      responseType: 'arraybuffer'
      // headers: {
      //   'Access-Control-Allow-Origin': '*',
      //   'Access-Control-Allow-Methods': 'GET, POST, PATCH, PUT, DELETE, OPTIONS',
      //   'Access-Control-Allow-Headers': 'Origin, Content-Type, X-Auth-Token',
      // }
    })
      .then(res => {
        const resp = res.data;
        this.setState({ config: domopool_pb.Config.deserializeBinary(resp).toObject() });
        console.log(this.state.config);
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

  render() {
    if (this.state.config.global != undefined) {
      if (this.state.config.global.displayStartup == undefined) {
        this.state.config.global.displayStartup = false;
      }
      this.state.ipDisabled = this.state.config.network.dhcp;
      return (
        <Formik
          enableReinitialize
          initialValues={{
            global: this.state.config.global,
            network: this.state.config.network,
            sensors: this.state.config.sensors,
          }
          }
          onSubmit={(values, actions) => {
            dataprovider.post('config')
              .then(res => {
                console.log(res.data)
              })
          }}
        >
          {props => (

            <Form>
              <GlobalFields {...props} />
              <br />
              <NetworkFields {...props} />
              <br />
              <Grid container spacing={2}>
                <Grid item>
                  <TwoutFields {...props} />
                </Grid>
                <Grid item>
                  <TambFields {...props} />
                </Grid>
                <Grid item>
                  <TwinFields {...props} />
                </Grid>
              </Grid>
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
