import {
    Field,
} from 'formik';
import {
    TextField,
    Checkbox,
} from 'formik-material-ui';
import FormControlLabel from '@material-ui/core/FormControlLabel';

// export const ConfigForm: React.SFC = () => {
const TwoutFields = (props: any) => {
    return (
        <div>
            <h3>Twin Values</h3>
            <FormControlLabel
                control={<Field
                    component={Checkbox}
                    name="sensors.twout.enabled"
                    type="checkbox"
                    color="primary"
                    disabled
                />}
                label="Enabled"
                labelPlacement="start"
            />
            <br />
            <FormControlLabel
                control={<Field
                    component={Checkbox}
                    name="sensors.twout.init"
                    type="checkbox"
                    color="primary"
                    disabled
                />}
                label="Init"
                labelPlacement="start"
            />
            <br />
            <Field
                component={TextField}
                name="sensors.twout.addr"
                type="text"
                label="Addr"
                disabled
            />
        </div >
    )
};

export default TwoutFields;
