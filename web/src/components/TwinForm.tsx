import * as React from "react";
import {
  Field,
} from 'formik';
import {
  TextField,
  Checkbox,
} from 'formik-material-ui';
import FormControlLabel from '@material-ui/core/FormControlLabel';

// export const ConfigForm: React.SFC = () => {
export default class TwinFields extends React.Component {
  constructor(
    public conf: any
  ) { super(conf); };
  // state = {
  //   tempDisabled: this.conf.values.network.dhcp,
  // }

  // public handleEnableClick = (event: React.ChangeEvent<HTMLInputElement>) => {
  //   this.state.tempDisabled = event.target.checked
  // }

  render() {
    return (
      <div>
        <h3>Twin Values</h3>
        <FormControlLabel
          control={<Field
            component={Checkbox}
            name="sensors.twin.enabled"
            type="checkbox"
            color="primary"
          />}
          label="Enabled"
          labelPlacement="start"
        />
        <br />
        <FormControlLabel
          control={<Field
            component={Checkbox}
            name="sensors.twin.init"
            type="checkbox"
            color="primary"
          />}
          label="Init"
          labelPlacement="start"
        />
        <br />
        <Field
          component={TextField}
          name="sensors.twin.addr"
          type="text"
          label="Addr"
          disabled
        />
      </div >
    )
  };
};
