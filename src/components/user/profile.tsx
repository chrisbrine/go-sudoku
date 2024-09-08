import { useState } from "react";
import { updateName, updateUsername } from "../../redux/actions";
import { useAppDispatch, useAppSelector } from "../../redux/hooks";
import { getName, getUserName } from "../../redux/selectors";

export default function User() {
  const dispatch = useAppDispatch();
  const onUpdateUserName = (username: string) =>
    dispatch(updateUsername(username));
  const onUpdateName = (name: string) => dispatch(updateName(name));
  const username = useAppSelector(getUserName);
  const name = useAppSelector(getName);

  const [newUsername, setNewUsername] = useState(username);
  const [newName, setNewName] = useState(name);

  const isValidUsername = () => {
    return newUsername.length > 0 && newName.length > 0;
  };

  return (
    <div className="profile">
      <div className="fields">
        <label>Username</label>
        <input
          type="text"
          placeholder="Username"
          value={newUsername}
          onChange={(e) => setNewUsername(e.target.value)}
        />
        <button
          className="action-button"
          disabled={!isValidUsername()}
          onClick={() => {
            onUpdateUserName(newUsername);
          }}
        >
          Update Username
        </button>
        <label>Name</label>
        <input
          type="text"
          placeholder="Name"
          value={newName}
          onChange={(e) => setNewName(e.target.value)}
        />
        <button
          className="action-button"
          disabled={!isValidUsername()}
          onClick={() => {
            onUpdateName(newName);
          }}
        >
          Update Name
        </button>
      </div>
    </div>
  );
}
