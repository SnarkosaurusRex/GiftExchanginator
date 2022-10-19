import React, { useState } from 'react';
import './App.scss';
import AddPersonForm, { PersonData } from './Components/AddPersonForm';

const App = () => {
    const [groupMembers, setGroupMembers] = useState<PersonData[]>([]);

    return (
        <div className="page-wrapper">
            <div className="top-content">
                <h1>Welcome to the Gift Exchanginator!</h1>
                <div className="instructions">Information about how this works/what to do with it/etc. will go here.</div>
            </div>

            <div className="add-person-section">
                <h2>{'Add a Group Member'}</h2>
                <AddPersonForm
                    onSubmit={(newPerson) => setGroupMembers(groupMembers.concat([newPerson]))}
                />
            </div>

            <div className="group-members-section">
                <h2>{'Group Members'}</h2>
                {'list of current group members will be displayed here'}
            </div>
            
            <div className="options-section">
                <h2>{'Options'}</h2>
                {'options will be displayed here'}
            </div>

            <button>{'Run'}</button>
        </div>
    );
}

export default App;
