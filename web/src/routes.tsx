import {
    Route,
    Switch
} from 'react-router-dom';
import ConfigView from './views/ConfigView';
import MetricsView from './views/MetricsView';
import AlarmsView from './views/AlarmsView';


export default function Routes() {


    return (
        <>
            <Switch>
                <Route exact path="/" component={MetricsView} />
                <Route exact path="/config" component={ConfigView} />
                <Route exact path="/alarms" component={AlarmsView} />
            </Switch>
        </>
    );
}
