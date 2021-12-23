import * as React from "react";
import dataprovider from '../dataprovider/dataprovider';
import CircularProgress from '@material-ui/core/CircularProgress';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Switch from '@material-ui/core/Switch';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormGroup from '@material-ui/core/FormGroup';

var domopool_pb = require('../proto/domopool_pb');

export default class MetricsView extends React.Component {
    constructor(
        props
    ) {
        super(props);
        this.state = {
            alarms: this.props,
        }
    };

    componentDidMount() {
        dataprovider.get(`/api/v1/alarms`, {
            responseType: 'arraybuffer'
            // headers: {
            //   'Access-Control-Allow-Origin': '*',
            //   'Access-Control-Allow-Methods': 'GET, POST, PATCH, PUT, DELETE, OPTIONS',
            //   'Access-Control-Allow-Headers': 'Origin, Content-Type, X-Auth-Token',
            // }
        })
            .then(res => {
                const resp = res.data;
                this.setState({ alarms: domopool_pb.Alarms.deserializeBinary(resp).toObject() });
                console.log(this.state.alarms);
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
        if (this.state.alarms.wpLow !== undefined) {
            // if (this.state.config.global.displayStartup == undefined) {
            //   this.state.config.global.displayStartup = false;
            // }

            return (
                <Card>
                    <CardContent>
                        <FormGroup>
                            <FormControlLabel
                                control={<Switch
                                    checked={this.state.alarms.wpLow}
                                    name="wpLow"
                                    inputProps={{ 'aria-label': 'secondary checkbox' }}
                                />}
                                label="Water pressure low"
                                labelPlacement="end"
                            />

                            <FormControlLabel
                                control={<Switch
                                    checked={this.state.alarms.wpHigh}
                                    name="wpHigh"
                                    inputProps={{ 'aria-label': 'secondary checkbox' }}
                                />}
                                label="Water pressure high"
                                labelPlacement="end"
                            />
                        </FormGroup>
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
