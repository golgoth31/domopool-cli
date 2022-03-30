import * as React from "react";
import dataprovider from '../dataprovider/dataprovider';
import {
    LinearProgress,
    Card,
    CardContent
} from "@mui/material";
import Chip from '@mui/material/Chip';
import AlarmOnIcon from '@mui/icons-material/AlarmOn';
import AlarmOffIcon from '@mui/icons-material/AlarmOff';

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
        })
            .then(res => {
                const resp = res.data;
                let config = domopool_pb.Config.deserializeBinary(resp).toObject();
                this.setState({ alarms: config.alarms });
                console.log(this.state.alarms);
            })
    }

    render() {
        if (this.state.alarms !== undefined) {
            return (
                <Card>
                    <CardContent>
                        <Chip
                            label="Water pressure low"
                            color={this.state.alarms.wpLow ? "error" : "success"}
                            icon={this.state.alarms.wpLow ? <AlarmOnIcon /> : <AlarmOffIcon />}
                        />
                        <Chip
                            label="Water pressure high"
                            color={this.state.alarms.wpHigh ? "error" : "success"}
                            icon={this.state.alarms.wpHigh ? <AlarmOnIcon /> : <AlarmOffIcon />}
                        />
                    </CardContent>
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
