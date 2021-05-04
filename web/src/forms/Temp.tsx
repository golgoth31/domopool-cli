import TempFields from './TempFields';
import {
    Grid,
} from "@material-ui/core";

const TempForm = (props: any) => {
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
