import { useEffect, useState } from "react";
import { DateTime } from "luxon";
import dataprovider from '../dataprovider/dataprovider';
import {
    LinearProgress,
    Card,
    CardContent,
    Grid,
    CardHeader,
    Typography,
} from "@mui/material";
import RelayButton from '../components/RelayButton';

var domopool_pb = require('../proto/domopool_pb');

export default function DashboardView() {
    const [config, setConfig] = useState(domopool_pb.Config);

    useEffect(
        () => {
            dataprovider.get(`/api/v1/config`, {
                responseType: 'arraybuffer'
            }).then(res => {
                const resp = res.data;
                setConfig(domopool_pb.Config.deserializeBinary(resp).toObject());
            })
        },
        [],
    );


    if (config !== undefined) {
        let date = DateTime.fromISO(
            config.metrics.time,
            {
                setZone: true,
            });
        let hours = date.hour.toString()
        if (date.hour < 10) {
            hours = "0" + date.hour
        }
        let minutes = date.minute.toString()
        if (date.minute < 10) {
            minutes = "0" + date.minute
        }

        let wp = config.metrics.wp / config.sensors.wp.vAccuracy;
        return (
            <Grid container direction="row" spacing={2} justifyContent="center" alignItems="center">
                <Grid item xs={6} >
                    <Card
                        variant="outlined"
                    >
                        <Grid container direction="row" spacing={0}>
                            <Grid item xs={6}>
                                <Card
                                    variant="outlined"
                                >
                                    <CardHeader
                                        title="Current Twater"
                                    />
                                    <CardContent>
                                        <Typography
                                            variant="h5"
                                            align="center"
                                        >
                                            {config.metrics.twater.toFixed(2) + " °C"}
                                        </Typography>
                                    </CardContent>
                                </Card>
                            </Grid>
                            <Grid item xs={6}>
                                <Card
                                    variant="outlined"
                                >
                                    <CardHeader
                                        title="Saved Twater"
                                    />
                                    <CardContent>
                                        <Typography
                                            variant="h5"
                                            align="center"
                                        >
                                            {config.metrics.savedTwater.toFixed(2) + " °C"}
                                        </Typography>
                                    </CardContent>
                                </Card>
                            </Grid>
                        </Grid >
                    </Card>
                </Grid>
                <Grid item xs={6}>
                    <Card
                        variant="outlined"
                    >
                        <CardHeader
                            title="Tamb"
                        />
                        <CardContent>
                            <Typography
                                variant="h5"
                                align="center"
                            >
                                {config.metrics.tamb.toFixed(2) + " °C"}
                            </Typography>
                        </CardContent>
                    </Card>
                </Grid>
                <Grid item xs={6}>
                    <Card
                        variant="outlined"
                    >
                        <CardHeader
                            title="Ph"
                        />
                        <CardContent>
                            <Typography
                                variant="h5"
                                align="center"
                            >
                                {config.metrics.ph}
                            </Typography>
                        </CardContent>
                    </Card>
                </Grid>
                <Grid item xs={6}>
                    <Card
                        variant="outlined"
                    >
                        <CardHeader
                            title="Water Pressure"
                        />
                        <CardContent>
                            <Typography
                                variant="h5"
                                align="center"
                            >
                                {wp.toFixed(2) + " Bar"}
                            </Typography>
                        </CardContent>
                    </Card>
                </Grid>

                <Grid item xs={12}>
                    <Card
                        variant="outlined"
                    >
                        <CardHeader
                            title="Infos"
                        />
                        <CardContent>
                            <Grid container direction="row" spacing={2} justifyContent="center" alignItems="center">
                                <Grid item xs={6} >
                                    <Grid container direction="column" spacing={2} alignItems="flex-start">
                                        <Grid item  >
                                            <Typography
                                                variant="h5"
                                            >
                                                {date.day + " " + date.monthShort + " " + date.year}
                                            </Typography>
                                        </Grid>
                                        <Grid item  >
                                            <Typography
                                                variant="h5"
                                            >
                                                {hours + "h" + minutes}
                                            </Typography>
                                        </Grid>
                                        <Grid item  >
                                            <Typography
                                                variant="h5"
                                            >
                                                {config.network.ip}
                                            </Typography>
                                        </Grid>
                                    </Grid>
                                </Grid>
                                <Grid item xs={6}>
                                    <Grid container direction="column" spacing={2} alignItems="flex-end">
                                        <Grid item  >
                                            <Typography
                                                variant="h5"
                                                color={config.states.netActive ? "success.main" : "error"}
                                            >
                                                NET
                                            </Typography>
                                        </Grid>
                                        <Grid item  >
                                            <Typography
                                                variant="h5"
                                                color={config.states.mqttConnected ? "success.main" : "error"}
                                            >
                                                MQTT
                                            </Typography>
                                        </Grid>
                                        <Grid item  >
                                            <Typography
                                                variant="h5"
                                                color={config.states.ntp ? "success.main" : "error"}
                                            >
                                                TIME
                                            </Typography>
                                        </Grid>
                                    </Grid>
                                </Grid>

                            </Grid>
                        </CardContent>
                    </Card>
                </Grid>
                <Grid item xs={12}>
                    <Card
                        variant="outlined"
                    >
                        <CardHeader
                            title="Relay"
                        />
                        <CardContent>
                            <Grid container direction="row" spacing={2} justifyContent="center" alignItems="center">
                                <Grid item xs={4} >
                                    <RelayButton relay='dyn' state={config.states.dynamic} offColor='primary' />
                                </Grid>
                                <Grid item xs={4} >
                                    <RelayButton relay='half' state={config.states.halfDay} offColor='primary' />
                                </Grid>
                                <Grid item xs={4} >
                                    <RelayButton relay='full' state={config.states.fullDay} offColor='primary' />
                                </Grid>
                                <Grid item xs={4} >
                                    <RelayButton relay='filter' state={config.states.filterOn} offColor='primary' />
                                </Grid>
                                <Grid item xs={4} >
                                    <RelayButton relay='ch' state={config.states.chOn} offColor='primary' />
                                </Grid>
                                <Grid item xs={4}>
                                    <RelayButton relay='ph' state={config.states.phOn} offColor='primary' />
                                </Grid>
                            </Grid>
                        </CardContent>
                    </Card>
                </Grid>
            </Grid >
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
