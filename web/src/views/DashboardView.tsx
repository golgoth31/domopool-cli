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
} from "@material-ui/core";
import useStyles from '../theme/useStyles';
import { makeStyles } from '@material-ui/core/styles';
import RelayButton from '../components/RelayButton';

var domopool_pb = require('../proto/domopool_pb');

const typographyStyles = makeStyles((theme) => ({
    text: {
        color: "#00ff00"
    }
}));

export default function DashboardView() {
    const [config, setConfig] = useState(domopool_pb.Config);
    const props = { state: false, offColor: '' };
    const classes = useStyles(props);
    const typoClasses = typographyStyles();

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
        return (
            <Grid container direction="row" spacing={2} justifyContent="center" alignItems="center">
                <Grid item xs={6} >
                    <Card
                        variant="outlined"
                    >
                        <CardHeader
                            title="Twater"
                        />
                        <CardContent>
                            <Typography
                                variant="h5"
                                align="center"
                            >
                                {config.metrics.twater + " °C"}
                            </Typography>
                        </CardContent>
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
                                {config.metrics.tamb + " °C"}
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
                                {config.metrics.wp + " Bar"}
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
                        <CardContent
                            className={classes.cardDashboard}
                        >
                            <Grid container direction="row" spacing={2} justifyContent="center" alignItems="center">
                                <Grid item xs={6} >
                                    <Grid container direction="column" spacing={2} alignItems="flex-start">
                                        <Grid item  >
                                            <Typography
                                                variant="h5"
                                            >
                                                {date.day + " " + date.monthLong + " " + date.year}
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
                                                className={typoClasses.text}
                                            // color={net_color}
                                            >
                                                NET
                                            </Typography>
                                        </Grid>
                                        <Grid item  >
                                            <Typography
                                                variant="h5"
                                                className={typoClasses.text}
                                            >
                                                MQTT
                                            </Typography>
                                        </Grid>
                                        <Grid item  >
                                            <Typography
                                                variant="h5"
                                                className={typoClasses.text}
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
                        <CardContent
                            className={classes.cardDashboard}
                        >
                            <Grid container direction="row" spacing={2} justifyContent="center" alignItems="center">
                                <Grid item xs={6} >
                                    <RelayButton relay='auto' state={config.states.automatic} offColor='#ff0000' />
                                </Grid>
                                <Grid item xs={6} >
                                    <RelayButton relay='recover' state={config.states.recover} offColor='#0000ff' />
                                </Grid>
                                <Grid item xs={4} >
                                    <RelayButton relay='filter' state={config.states.filterOn} offColor='#0000ff' />
                                </Grid>
                                <Grid item xs={4} >
                                    <RelayButton relay='ch' state={config.states.chOn} offColor='#0000ff' />
                                </Grid>
                                <Grid item xs={4}>
                                    <RelayButton relay='ph' state={config.states.phOn} offColor='#0000ff' />
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
