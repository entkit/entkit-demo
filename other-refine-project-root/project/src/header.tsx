// Code generated by EntKit. DO NOT EDIT.
// ---------------------------------------------------------
// 
// Copyright (C) 2023 EntKit. All Rights Reserved.
// 
// This code is part of the EntKit library and is generated
// automatically to ensure optimal functionality and maintainability.
// Any changes made directly to this file may be overwritten
// by future code generation, leading to unexpected behavior.
// 
// Please refer to the EntKit documentation for instructions on
// how to modify the library, extend its functionality or contribute
// to the project: https://entkit.com
// ---------------------------------------------------------

import {Col, Grid, Layout as AntdLayout, Row} from "antd";
import {SearchComponent} from "search-component";

const {Header: AntdHeader} = AntdLayout;

const {useBreakpoint} = Grid;


interface IOptionGroup {
    value: string;
    label: string | React.ReactNode;
}

interface IOptions {
    label: string | React.ReactNode;
    options: IOptionGroup[];
}

export const Header: React.FC = () => {
    const screens = useBreakpoint();

    return (
        <AntdHeader
            style={
                {
                    padding: "0 24px",
                    background: "white",
                }
            }
        >
            <Row
                align="middle"
                style={
                    {
                        justifyContent: screens.sm ? "space-between" : "end",
                    }
                }
            >
                <Col xs={0} sm={12}>
                    <SearchComponent/>
                </Col>
                <Col>

                </Col>
            </Row>
        </AntdHeader>
    );
};