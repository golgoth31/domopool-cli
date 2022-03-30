import TempFields from './TempFields';
import {
    Grid,
} from "@mui/material";

const TempForm = (props) => {
    return (
        <Grid container spacing={2}>
            <Grid item>
                <TempFields sensors={props.config} name="twout" />
            </Grid>
            <Grid item>
                <TempFields sensors={props.config} name="tamb" />
            </Grid>
            <Grid item>
                <TempFields sensors={props.config} name="twin" />
            </Grid>
        </Grid>
    )
};

export default TempForm;
