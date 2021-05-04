import {
    Field,
} from 'formik';
import {
    TextField,
    Checkbox,
} from 'formik-material-ui';
import FormControlLabel from '@material-ui/core/FormControlLabel';

// export const ConfigForm: React.SFC = () => {
const TwinFields = (props: any) => {
    return (
        <div>
            <h3>Twin Values</h3>
            <FormControlLabel
                control={<Field
                    component={Checkbox}
                    name="sensors.twin.enabled"
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
                    name="sensors.twin.init"
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
                name="sensors.twin.addr"
                type="text"
                label="Addr"
                disabled
            />
        </div >
    )
};

export default TwinFields;
