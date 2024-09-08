import { useAppDispatch, useAppSelector } from "../../../redux/hooks";
import { setPickHint as onSetPickHint } from "../../../redux/actions";
import "./pickType.css";
import { getPickHint } from "../../../redux/selectors";

export default function PickType() {
  const dispatch = useAppDispatch();
  const setPickHint = (pickHint: boolean) => dispatch(onSetPickHint(pickHint));

  const pickHint = useAppSelector(getPickHint);
  /* Handle as buttons for 'Hints' or 'Move' */
  return (
    <div className="pick-type">
      <button
        className={`pick-hint ${pickHint ? "active" : "inactive"}`}
        onClick={() => setPickHint(true)}
      >
        Hints
      </button>
      <button
        className={`pick-move ${!pickHint ? "active" : "inactive"}`}
        onClick={() => setPickHint(false)}
      >
        Move
      </button>
    </div>
  );
}
