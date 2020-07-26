import * as React from "react";
import {
  Field,
} from 'formik';
import {
  TextField,
  Checkbox,
} from 'formik-material-ui';
import { Grid } from "@material-ui/core";
import FormControlLabel from '@material-ui/core/FormControlLabel';
import isIp from 'is-ip';

// export const ConfigForm: React.SFC = () => {
export default class GlobalFields extends React.Component {


  private validateNumber(value: any) {
    let error;
    if (!value) {
      error = 'Required';
    } else if (!/^[0-9]+$/i.test(value)) {
      error = 'Invalid duration';
    }
    return error;
  }

  render() {
    return (
      <div>
        <h3>Global values</h3>
        <Grid container spacing={2}>
          <Grid item>
            <Field
              component={TextField}
              name="global.lcdbacklightduration"
              type="text"
              label="lcdBacklightDuration"
              validate={this.validateNumber}
            />
          </Grid>
          <Grid item>
            <Field
              component={TextField}
              name="global.ackduration"
              type="text"
              label="ackDuration"
              validate={this.validateNumber}
            />
          </Grid>
          <Grid item>
            <Field
              component={TextField}
              name="global.acktone"
              type="text"
              label="ackTone"
              validate={this.validateNumber}
            />
          </Grid>
          <Grid item>
            <FormControlLabel
              control={<Field
                component={Checkbox}
                name="global.serialout"
                type="checkbox"
                color="primary"
              // onClick={this.handleDHCPClick}
              />}
              label="Print to Serial"
              labelPlacement="start"
            />
          </Grid>
          <Grid item>
            <FormControlLabel
              control={<Field
                component={Checkbox}
                name="global.displaystartup"
                type="checkbox"
                color="primary"
              // onClick={this.handleDHCPClick}
              />}
              label="Display startup data"
              labelPlacement="start"
            />
          </Grid>
        </Grid>
      </div >
    )
  };
};
