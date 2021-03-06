<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: consensus.proto

namespace Protocol;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>protocol.PbftViewChangeWithRawValue</code>
 */
class PbftViewChangeWithRawValue extends \Google\Protobuf\Internal\Message
{
    /**
     * view change env
     *
     * Generated from protobuf field <code>.protocol.PbftEnv view_change_env = 1;</code>
     */
    private $view_change_env = null;
    /**
     *prepared messages large than n
     *
     * Generated from protobuf field <code>.protocol.PbftPreparedSet prepared_set = 2;</code>
     */
    private $prepared_set = null;

    public function __construct() {
        \GPBMetadata\Consensus::initOnce();
        parent::__construct();
    }

    /**
     * view change env
     *
     * Generated from protobuf field <code>.protocol.PbftEnv view_change_env = 1;</code>
     * @return \Protocol\PbftEnv
     */
    public function getViewChangeEnv()
    {
        return $this->view_change_env;
    }

    /**
     * view change env
     *
     * Generated from protobuf field <code>.protocol.PbftEnv view_change_env = 1;</code>
     * @param \Protocol\PbftEnv $var
     * @return $this
     */
    public function setViewChangeEnv($var)
    {
        GPBUtil::checkMessage($var, \Protocol\PbftEnv::class);
        $this->view_change_env = $var;

        return $this;
    }

    /**
     *prepared messages large than n
     *
     * Generated from protobuf field <code>.protocol.PbftPreparedSet prepared_set = 2;</code>
     * @return \Protocol\PbftPreparedSet
     */
    public function getPreparedSet()
    {
        return $this->prepared_set;
    }

    /**
     *prepared messages large than n
     *
     * Generated from protobuf field <code>.protocol.PbftPreparedSet prepared_set = 2;</code>
     * @param \Protocol\PbftPreparedSet $var
     * @return $this
     */
    public function setPreparedSet($var)
    {
        GPBUtil::checkMessage($var, \Protocol\PbftPreparedSet::class);
        $this->prepared_set = $var;

        return $this;
    }

}

