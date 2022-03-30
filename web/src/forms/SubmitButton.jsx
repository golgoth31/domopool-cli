import * as React from "react";
import {
    makeStyles,
    createStyles
} from '@mui/styles';
import Button from '@mui/material/Button';

const useStyles = makeStyles((theme) =>
    createStyles({
        root: {
            '& > *': {
                margin: theme.spacing(1),
            },
        },
        input: {
            display: 'none',
        },
    }),
);

export default function SubmitButton() {
    const classes = useStyles();

    return (
        <div>
            <input
                className={classes.input}
                id="contained-button-file"
                type="submit"
            />
            <label htmlFor="contained-button-file">
                <Button variant="contained" color="primary" component="span">
                    Submit
                </Button>
            </label>
        </div>
    );
}
