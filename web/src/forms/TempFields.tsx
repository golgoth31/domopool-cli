import {
    TextField,
    Checkbox,
} from '@material-ui/core';
import FormControlLabel from '@material-ui/core/FormControlLabel';

// export const ConfigForm: React.SFC = () => {
const TempFields = (props: any) => {
    console.debug(props);
    return (
        <div>
            <h3>{props.name} Values</h3>
            <FormControlLabel
                control={<Checkbox
                    color="primary"
                    checked={props.sensors[props.name].enabled}
                    disabled
                />}
                label="Enabled"
            // labelPlacement="start"
            />
            <br />
            <FormControlLabel
                control={<Checkbox
                    color="primary"
                    checked={props.sensors[props.name].init}
                    disabled
                />}
                label="Init"
            // labelPlacement="start"
            />
            <br />
            <TextField
                value={props.sensors[props.name].addr}
                type="text"
                label="Addr"
                disabled
            />
        </div >
    )
};

export default TempFields;
