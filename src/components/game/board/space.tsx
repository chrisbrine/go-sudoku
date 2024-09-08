import { setHint, setHintRemove, setMove } from "../../../redux/actions";
import "./space.css";
import { IRowColNum } from "../../../types";
import React from "react";
import { store } from "../../../redux/store";
import { RootState } from "../../../redux/types";
import { connect, ConnectedProps } from "react-redux";

interface SpaceProps extends PropsFromRedux {
  row: number;
  col: number;
}

export class Space extends React.Component<SpaceProps> {
  state = {
    number: this.props.board[this.props.col][this.props.row],
    hints: this.props.hints[this.props.col][this.props.row],
  };

  onMove = (num: number) => {
    const { row, col, setMove } = this.props;
    setMove({ row, col, num });
  };

  onHint = (num: number) => {
    const { row, col, setHint } = this.props;
    setHint({ row, col, num });
  };

  onHintRemove = (num: number) => {
    const { row, col, setHintRemove } = this.props;
    setHintRemove({ row, col, num });
  };

  async componentDidMount() {
    store.subscribe(() => {
      // wait one seconds, then force update
      this.setState({
        number:
          store.getState().game.game.Board[this.props.col][this.props.row],
        hints: store.getState().game.game.Hints[this.props.col][this.props.row],
      });
      this.forceUpdate();
    });
  }

  render() {
    const { pickHint, numbersLeft } = this.props;
    const { number, hints } = this.state;

    const classes = ["number"];
    if (number > 0) {
      classes.push("correct");
      classes.push(`number-${number}`);
      if (numbersLeft[number - 1] === 0) {
        classes.push("completed-number");
      } else {
        classes.push("incomplete-number");
      }
    } else {
      classes.push("todo");
      classes.push(`no-number`);
    }

    return (
      <div className="space">
        <div className={classes.join(" ")}>{number === 0 ? "" : number}</div>
        {number === 0 ? (
          <div className="hints">
            {Object.values(hints).map((h, i) => (
              <div
                key={i}
                className={`hint ${h ? "active" : "inactive"} hint-${i + 1}`}
                onClick={() =>
                  pickHint
                    ? h
                      ? this.onHintRemove(i + 1)
                      : this.onHint(i + 1)
                    : this.onMove(i + 1)
                }
              >
                {i + 1}
              </div>
            ))}
          </div>
        ) : null}
      </div>
    );
  }
}

const mapStateToProps = (state: RootState) => {
  return {
    board: state.game.game.Board,
    pickHint: state.game.pickHint,
    hints: state.game.game.Hints,
    numbersLeft: state.game.game.NumbersLeft,
  };
};

const mapDispatchToProps = (dispatch: any) => {
  return {
    setHint: ({ row, col, num }: IRowColNum) =>
      dispatch(setHint({ row: row + 1, col: col + 1, num })),
    setHintRemove: ({ row, col, num }: IRowColNum) =>
      dispatch(setHintRemove({ row: row + 1, col: col + 1, num })),
    setMove: ({ row, col, num }: IRowColNum) =>
      dispatch(setMove({ row: row + 1, col: col + 1, num })),
  };
};

export const connector = connect(mapStateToProps, mapDispatchToProps);

type PropsFromRedux = ConnectedProps<typeof connector>;

export default connector(Space);
