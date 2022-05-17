import * as React from "react";
import {
    Button,
    Dialog,
    DialogTitle,
    Grid,
    Slider,
    Input,
    DialogContent,
    DialogActions,
    Switch,
    Typography,
} from "@mui/material";
import {
    makeStyles
} from '@mui/styles';
import dataprovider from '../dataprovider/dataprovider';

var domopool_pb = require('../proto/domopool_pb');

const useStyles = makeStyles({
    input: {
        width: 42,
    },
});

const relayButtonStyles = makeStyles((theme) => ({
    relayButton: {
        justifyContent: 'center',
        width: '100%',
        alignContent: 'center',
    },
}));

function apiPath(relay) {
    // Set api path
    let apiPath = '/api/v1/';
    switch (relay) {
        case 'dyn':
        case 'half':
        case 'full':
            apiPath = apiPath + 'pump/time';
            break;
        case 'filter':
        case 'ch':
        case 'ph':
            apiPath = apiPath + 'relay'
            break;
    }

    return apiPath
}

function SimpleDialog(props) {
    const classes = useStyles();
    // const { onClose, duration, relay, state, open } = props;
    const [value, setValue] = React.useState(0);
    const [state, setState] = React.useState(props.state);

    const handleClose = () => {
        props.onClose(props.duration);
    };

    const handleSet = () => {
        // Relay object config
        let configRelay = new domopool_pb.Relay();
        configRelay.setDuration(value);
        if (state) {
            configRelay.setState(domopool_pb.Relay_states["START"]);
        } else {
            configRelay.setState(domopool_pb.Relay_states["STOP"]);
            configRelay.setDuration(0);
        }
        configRelay.setRelay(domopool_pb.Relay_names[props.relay.toUpperCase()]);
        const relayBuffer = configRelay.serializeBinary();

        dataprovider.post(
            apiPath(props.relay),
            relayBuffer,
            {
                responseType: 'arraybuffer',
                headers: { 'Content-Type': 'application/octet-stream' }
            }
        ).then(function (response) {
            console.log(response)
        }).catch(function (response) {
            console.log(response)
        })
        props.onClose(props.duration);
    };

    const handleSliderChange = (event, newValue) => {
        setValue(newValue);
    };

    const handleInputChange = (event) => {
        setValue(event.target.value === '' ? '' : Number(event.target.value));
    };

    const handleBlur = () => {
        if (value < 0) {
            setValue(0);
        } else if (value > 300) {
            setValue(300);
        }
    };

    const handleSwitch = (event) => {
        setState(event.target.checked);
    };

    return (
        <Dialog onClose={handleClose} aria-labelledby="simple-dialog-title" open={props.open}>
            <DialogTitle id="simple-dialog-title">Set duration for "{props.relay}" (in minutes)</DialogTitle>
            <DialogContent>
                <Grid container spacing={2} alignItems="center">
                    <Grid item xs={4}>
                        State
                    </Grid>
                    <Grid item xs={8}>
                        <Switch
                            color="primary"
                            checked={state}
                            onChange={handleSwitch}
                        />
                    </Grid>
                    <Grid item xs={3}>
                        Duration
                    </Grid>
                    <Grid item xs={6}>
                        <Slider
                            value={typeof value === 'number' ? value : 0}
                            onChange={handleSliderChange}
                            aria-labelledby="input-slider"
                            step={10}
                            max={300}
                            disabled={!state}
                        />
                    </Grid>
                    <Grid item xs={3}>
                        <Input
                            className={classes.input}
                            value={value}
                            margin="dense"
                            onChange={handleInputChange}
                            onBlur={handleBlur}
                            inputProps={{
                                step: 10,
                                min: 0,
                                max: 300,
                                type: 'number',
                                'aria-labelledby': 'input-slider',
                            }}
                            disabled={!state}
                        />
                    </Grid>
                </Grid>
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClose} color="primary">
                    Cancel
                </Button>
                <Button onClick={handleSet} color="primary" autoFocus>
                    Set
                </Button>
            </DialogActions>
        </Dialog >
    );
}

export default function RelayButton(props) {
    const [open, setOpen] = React.useState(false);
    const [duration, setDuration] = React.useState(0);
    const handleClickOpen = () => {
        switch (props.relay) {
            case 'filter':
            case 'ch':
            case 'ph':
                setOpen(true);
                break;
            case 'dyn':
            case 'half':
            case 'full':
                dataprovider.post(
                    apiPath(props.relay)
                ).then(function (response) {
                    console.log(response)
                }).catch(function (response) {
                    console.log(response)
                })
        }
    };

    const handleClose = (value) => {
        setOpen(false);
        setDuration(value);
    };
    const classes = relayButtonStyles(props);

    return (
        <div>
            <Button
                variant="contained"
                className={classes.relayButton}
                onClick={handleClickOpen}
                color={props.state ? "success" : props.offColor}
            >
                <Typography
                    variant="h5"
                >
                    {props.relay}
                </Typography>
            </Button>
            <SimpleDialog
                duration={duration}
                relay={props.relay}
                state={props.state}
                open={open}
                onClose={handleClose}
            />
        </div >
    );
};
