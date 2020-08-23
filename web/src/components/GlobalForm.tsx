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
              name="global.lcdBacklightDuration"
              type="text"
              label="lcdBacklightDuration"
              validate={this.validateNumber}
            />
          </Grid>
          <Grid item>
            <Field
              component={TextField}
              name="global.ackDuration"
              type="text"
              label="ackDuration"
              validate={this.validateNumber}
            />
          </Grid>
          <Grid item>
            <Field
              component={TextField}
              name="global.ackTone"
              type="text"
              label="ackTone"
              validate={this.validateNumber}
            />
          </Grid>
          <Grid item>
            <FormControlLabel
              control={<Field
                component={Checkbox}
                name="global.serialOut"
                type="checkbox"
                color="primary"
              />}
              label="Print to Serial"
              labelPlacement="start"
            />
          </Grid>
          <Grid item>
            <FormControlLabel
              control={<Field
                component={Checkbox}
                name="global.displayStartup"
                type="checkbox"
                color="primary"
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
