import * as React from "react";
import {
    Field,
} from 'formik';
import {
    TextField,
    Checkbox,
} from 'formik-mui';
import isIp from 'is-ip';
import {
    Grid,
    FormControlLabel
} from "@mui/material";

// export const ConfigForm: React.SFC = () => {
export default class NetworkForm extends React.Component {
    constructor(
        props
    ) {
        super(props);
        this.state = {
            ipDisabled: this.props.values.network.dhcp,
        }
    };

    validateIP = (value) => {
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
    handleDHCPClick = (event) => {
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
