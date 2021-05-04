import * as React from "react";
import {
    Formik,
    Form,
} from 'formik';
import TwoutFields from '../components/TwoutForm';
import TwinFields from '../components/TwinForm';
import TambFields from '../components/TambForm';
import {
    Grid,
} from "@material-ui/core";

const TempForm = (props: any) => {
    return (
        <Formik
            enableReinitialize
            initialValues={{
                sensors: props.config,
            }
            }
            onSubmit={() => {
            }}
        >
            {props => (
                <Form>
                    <Grid container spacing={2}>
                        <Grid item>
                            <TwoutFields {...props} />
                        </Grid>
                        <Grid item>
                            <TambFields {...props} />
                        </Grid>
                        <Grid item>
                            <TwinFields {...props} />
                        </Grid>
                    </Grid>
                </Form>

            )}
        </Formik >
    )
};

export default TempForm;
