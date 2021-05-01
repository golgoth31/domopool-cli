import * as React from "react";
import {
    Field,
} from 'formik';
import {
    TextField,
    Checkbox,
} from 'formik-material-ui';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import isIp from 'is-ip';
import { Grid } from "@material-ui/core";

// export const ConfigForm: React.SFC = () => {
export default class NetworkFields extends React.Component {
    constructor(
        private conf: any
    ) { super(conf); };
    state = {
        ipDisabled: this.conf.values.network.dhcp,
    }

    private validateIP = (value: any) => {
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
    public handleDHCPClick = (event: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({ ipDisabled: event.target.checked });
    }

    render() {
        return (
            <div>
                <h3>Network values</h3>
                <FormControlLabel
                    control={<Field
                        component={Checkbox}
                        name="network.dhcp"
                        type="checkbox"
                        color="primary"
                        onClick={this.handleDHCPClick}
                    />}
                    label="Use DHCP"
                    labelPlacement="start"
                />
                <br />
                <Grid container spacing={2}>
                    <Grid item>
                        <Field
                            component={TextField}
                            name="network.ip"
                            type="text"
                            label="IP address"
                            validate={this.validateIP}
                            disabled={this.state.ipDisabled}
                        />
                    </Grid>
                    <Grid item>
                        <Field
                            component={TextField}
                            name="network.netmask"
                            type="text"
                            label="Netmask"
                            validate={this.validateIP}
                            disabled={this.state.ipDisabled}
                        />
                    </Grid>
                    <Grid item>
                        <Field
                            component={TextField}
                            name="network.gateway"
                            type="text"
                            label="Gateway"
                            validate={this.validateIP}
                            disabled={this.state.ipDisabled}
                        />
                    </Grid>
                    <Grid item>
                        <Field
                            component={TextField}
                            name="network.dns"
                            type="text"
                            label="DNS server IP"
                            validate={this.validateIP}
                            disabled={this.state.ipDisabled}
                        />
                    </Grid>
                </Grid>
            </div >
        )
    };
};
