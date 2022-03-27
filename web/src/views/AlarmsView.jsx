import * as React from "react";
import dataprovider from '../dataprovider/dataprovider';
import {
    LinearProgress,
    Card,
    CardContent,
    Switch,
    FormControlLabel,
    FormGroup,
} from "@material-ui/core";

var domopool_pb = require('../proto/domopool_pb');

export default class AlarmsView extends React.Component {
    constructor(
        alarms
    ) {
        super(alarms);
        this.state = {
            alarms: this.alarms,
        }
    };

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
                let config = domopool_pb.Config.deserializeBinary(resp).toObject();
                this.setState({ alarms: config.alarms });
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
        if (this.state.alarms !== undefined) {
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
                <Card>
                    <CardContent>
                        <LinearProgress />
                    </CardContent>
                </Card>
            )
        }
    };
};
