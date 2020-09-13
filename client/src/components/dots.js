import React from "react";
import Dot from "./dot";
import styles from "./styles.module.css";

const Dots = ({ data }) => {
  const dots = data.map((datum, index) => {
    return <Dot key={`data-${index}-${datum}`} int={datum} />;
  });

  return <div className={styles.container}>{dots}</div>;
};

export default Dots;
