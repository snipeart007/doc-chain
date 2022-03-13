import React from "react";
import clsx from "clsx";
import styles from "./styles.module.css";

type FeatureItem = {
    title: string;
    Svg: React.ComponentType<React.ComponentProps<"svg">>;
    description: JSX.Element;
};

const FeatureList: FeatureItem[] = [
    {
        title: "Fully Secured",
        Svg: require("@site/static/img/undraw_docusaurus_mountain.svg").default,
        description: (
            <>
                DocChain is going to be implemented with proper security
                measures and will use Go, a compiled and type-safe language
                ensuring safety.
            </>
        ),
    },
    {
        title: "Trusted Cryptography",
        Svg: require("@site/static/img/undraw_docusaurus_tree.svg").default,
        description: (
            <>
                DocChain secures the blockchain with trusted crypto algorithms
                like SHA-256, RSA, etc.
            </>
        ),
    },
    {
        title: "Single Governing Authority",
        Svg: require("@site/static/img/undraw_docusaurus_react.svg").default,
        description: (
            <>
                DocChain is designed to be controlled under a single main
                system. This ensures that no one can ever alter with the data.
            </>
        ),
    },
];

function Feature({ title, Svg, description }: FeatureItem) {
    return (
        <div className={clsx("col col--4")}>
            <div className="text--center">
                <Svg className={styles.featureSvg} role="img" />
            </div>
            <div className="text--center padding-horiz--md">
                <h3>{title}</h3>
                <p>{description}</p>
            </div>
        </div>
    );
}

export default function HomepageFeatures(): JSX.Element {
    return (
        <section className={styles.features}>
            <div className="container">
                <div className="row">
                    {FeatureList.map((props, idx) => (
                        <Feature key={idx} {...props} />
                    ))}
                </div>
            </div>
        </section>
    );
}
