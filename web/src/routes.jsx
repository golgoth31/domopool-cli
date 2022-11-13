import {
    Route,
    Routes
} from 'react-router-dom';
import ConfigView from './views/ConfigView';
import DashboardView from './views/DashboardView';
import AlarmsView from './views/AlarmsView';


export default function AppRoutes() {


    return (
        <>
            <Routes>
                <Route exact path="/" element={<DashboardView />} />
                <Route exact path="/config" element={<ConfigView />} />
                <Route exact path="/alarms" element={<AlarmsView />} />
            </Routes>
        </>
    );
}
