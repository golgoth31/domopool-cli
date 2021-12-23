import {
    Route,
    Switch
} from 'react-router-dom';
import ConfigView from './views/ConfigView';
import DashboardView from './views/DashboardView';
import AlarmsView from './views/AlarmsView';


export default function Routes() {


    return (
        <>
            <Switch>
                <Route exact path="/" component={DashboardView} />
                <Route exact path="/config" component={ConfigView} />
                <Route exact path="/alarms" component={AlarmsView} />
            </Switch>
        </>
    );
}
