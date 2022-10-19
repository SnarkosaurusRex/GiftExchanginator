import React, { useState } from 'react';
import TextInput from './TextInput';

export interface PersonData {
    name: string;
    email: string;
    roommate: string;
    significantOther: string;
}

interface Props {
    onSubmit: (personData: PersonData) => void;
}

const AddPersonForm = (props: Props) => {
    const [personData, setPersonData] = useState<PersonData>({
        name: '',
        email: '',
        roommate: '',
        significantOther: '',
    });
    const { onSubmit } = props;

    return (
        <div>
            <form onSubmit={(e) => {
                e.preventDefault();
                onSubmit(personData);
            }}>
                <TextInput
                    name="name"
                    label="Name"
                    value={personData.name}
                    onChange={(name) => setPersonData({...personData, name})}
                    required={true}
                />
                <TextInput
                    name="emailAddress"
                    label="Email Address"
                    value={personData.email}
                    onChange={(email) => setPersonData({...personData, email})}
                    required={true}
                />
                <TextInput
                    name="roommate"
                    label="Roommate"
                    value={personData.roommate}
                    onChange={(roommate) => setPersonData({...personData, roommate})}
                />
                <TextInput
                    name="significantOther"
                    label="Significant Other"
                    value={personData.significantOther}
                    onChange={(significantOther) => setPersonData({...personData, significantOther})}
                />

                <button type="submit">{'Add Person'}</button>
            </form>
        </div>
    );
}

export default AddPersonForm;