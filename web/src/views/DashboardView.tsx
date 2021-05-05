import { useEffect, useState } from "react";
import dataprovider from '../dataprovider/dataprovider';
import {
    LinearProgress,
    Card,
    CardContent,
    Grid,
    TextField,
} from "@material-ui/core";
// import useStyles from '../theme/useStyles';

var domopool_pb = require('../proto/domopool_pb');

export default function DashboardView() {
    const [config, setConfig] = useState(domopool_pb.Config);
    // const classes = useStyles();

    useEffect(
        () => {
            dataprovider.get(`/api/v1/config`, {
                responseType: 'arraybuffer'
            })
                .then(res => {
                    const resp = res.data;
                    setConfig(domopool_pb.Config.deserializeBinary(resp).toObject());
                })
        },
        [],
    );


    if (config !== undefined) {
        return (
            <Card>
                <CardContent>
                    <Grid container spacing={2} justify="center">
                        <Grid item xs={12}>
                            <Grid container justify="center" spacing={2}>
                                <Grid key={1} item>
                                    <TextField
                                        label="Twater"
                                        defaultValue={config.metrics.twater + " °C"}
                                        InputProps={{
                                            readOnly: true,
                                        }}
                                        variant="outlined"
                                    />
                                </Grid>
                                <Grid key={2} item>
                                    <TextField
                                        label="Tamb"
                                        defaultValue={config.metrics.tamb + " °C"}
                                        InputProps={{
                                            readOnly: true,
                                        }}
                                        variant="outlined"
                                    />
                                </Grid>
                            </Grid>
                            <Grid container justify="center" spacing={2}>
                                <Grid key={1} item>
                                    <TextField
                                        label="Ph"
                                        defaultValue={config.metrics.ph}
                                        InputProps={{
                                            readOnly: true,
                                        }}
                                        variant="outlined"
                                    />
                                </Grid>
                                <Grid key={2} item>
                                    <TextField
                                        label="Water Pressure"
                                        defaultValue={config.metrics.wp + " Bar"}
                                        InputProps={{
                                            readOnly: true,
                                        }}
                                        variant="outlined"
                                    />
                                </Grid>
                            </Grid>
                        </Grid>
                    </Grid>
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
