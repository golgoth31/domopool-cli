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
        <Field
          component={TextField}
          name="global.lcdbacklightduration"
          type="text"
          label="lcdBacklightDuration"
          validate={this.validateNumber}
        />
      </div >
    )
  };
};
