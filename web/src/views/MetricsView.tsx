import * as React from "react";
import dataprovider from '../dataprovider/dataprovider';
import {
  Formik,
  Form,
} from 'formik';
import SubmitButton from "../components/SubmitButton";
import CircularProgress from '@material-ui/core/CircularProgress';
import NetworkFields from '../components/NetworkForm';
import GlobalFields from '../components/GlobalForm';
import TwoutFields from '../components/TwoutForm';
import TwinFields from '../components/TwinForm';
import TambFields from '../components/TambForm';
import { Grid } from "@material-ui/core";
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';

var domopool_pb = require('../proto/domopool_pb');
// import { } from './../proto/domopool_pb';

export default class MetricsView extends React.Component {
  constructor(
    private config: any
  ) { super(config); };

  state = {
    config: this.config,
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
    if (this.state.config.metrics != undefined) {
      // if (this.state.config.global.displayStartup == undefined) {
      //   this.state.config.global.displayStartup = false;
      // }

      return (
        <Card>
          <CardContent>
            {this.state.config.metrics.twater}
          </CardContent>
          {/* <CardActions>
            <Button size="small">Learn More</Button>
          </CardActions> */}
        </Card>
      )
    } else {
      return (
        <CircularProgress />
      )
    }
  };
};
