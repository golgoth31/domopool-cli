import * as React from "react";
import dataprovider from '../dataprovider/dataprovider';
import {
    Formik,
    Form,
} from 'formik';
import SubmitButton from "../components/SubmitButton";
import NetworkFields from '../components/NetworkForm';
import GlobalFields from '../components/GlobalForm';
import TwoutFields from '../components/TwoutForm';
import TwinFields from '../components/TwinForm';
import TambFields from '../components/TambForm';
import {
    Grid,
    LinearProgress,
    Card,
    CardContent,
} from "@material-ui/core";

var domopool_pb = require('../proto/domopool_pb');

export default class ConfigView extends React.Component {
    constructor(
        private conf: any
    ) { super(conf); };

    state = {
        config: this.conf,
        ipDisabled: false,
    }

    componentDidMount() {
        dataprovider.get(`/api/v1/config`, {
            responseType: 'arraybuffer'
        })
            .then(res => {
                const resp = res.data;
                this.setState({ config: domopool_pb.Config.deserializeBinary(resp).toObject() });
                console.log(this.state.config);
                this.setState({ ipDisabled: this.state.config.network.dhcp });
            });
    }

    render() {
        if (this.state.config.global !== undefined) {
            if (this.state.config.global.displayStartup === undefined) {
                this.setState({ config: { global: { displayStartup: false } } });
            }

            return (
                <Card>
                    <CardContent>
                        <Formik
                            enableReinitialize
                            initialValues={{
                                global: this.state.config.global,
                            }
                            }
                            onSubmit={(values, actions) => {
                                dataprovider.post(this.state.config.toArrayBuffer(),
                                    {
                                        responseType: 'arraybuffer',
                                        headers: { 'Content-Type': 'application/octet-stream' }
                                    }
                                ).then(function (response) {
                                    console.log(response)
                                })
                                    .catch(function (response) {
                                        console.log(response)
                                    })
                            }}
                        >
                            {props => (

                                <Form>
                                    <GlobalFields {...props} />
                                    <br />
                                    <SubmitButton />
                                </Form>

                            )}
                        </Formik >
                        <Formik
                            enableReinitialize
                            initialValues={{
                                network: this.state.config.network,
                            }
                            }
                            onSubmit={(values, actions) => {
                                dataprovider.post(this.state.config.toArrayBuffer(),
                                    {
                                        responseType: 'arraybuffer',
                                        headers: { 'Content-Type': 'application/octet-stream' }
                                    }
                                ).then(function (response) {
                                    console.log(response)
                                })
                                    .catch(function (response) {
                                        console.log(response)
                                    })
                            }}
                        >
                            {props => (

                                <Form>
                                    <NetworkFields {...props} />
                                    <br />
                                    <SubmitButton />
                                </Form>

                            )}
                        </Formik >
                        <Formik
                            enableReinitialize
                            initialValues={{
                                sensors: this.state.config.sensors,
                            }
                            }
                            onSubmit={(values, actions) => {
                                dataprovider.post(this.state.config.toArrayBuffer(),
                                    {
                                        responseType: 'arraybuffer',
                                        headers: { 'Content-Type': 'application/octet-stream' }
                                    }
                                ).then(function (response) {
                                    console.log(response)
                                })
                                    .catch(function (response) {
                                        console.log(response)
                                    })
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
                                    <br />
                                    <SubmitButton />
                                </Form>

                            )}
                        </Formik >
                    </CardContent>
                </Card>
            )
        } else {
            return (
                <Card>
                    <CardContent>
                        <LinearProgress />
                    </CardContent>
                </Card>
            )
        }
    };
};
