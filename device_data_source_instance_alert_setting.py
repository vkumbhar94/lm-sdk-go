# coding: utf-8

"""
    LogicMonitor REST API

    LogicMonitor is a SaaS-based performance monitoring platform that provides full visibility into complex, hybrid infrastructures, offering granular performance monitoring and actionable data and insights. logicmonitor_sdk enables you to manage your LogicMonitor account programmatically.  # noqa: E501

    OpenAPI spec version: 1.0.0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from logicmonitor_sdk.models.device_group_alert_threshold_info import DeviceGroupAlertThresholdInfo  # noqa: F401,E501


class DeviceDataSourceInstanceAlertSetting(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'global_alert_expr': 'str',
        'alert_clear_interval': 'int',
        'disable_alerting': 'bool',
        'alert_expr_note': 'str',
        'data_point_description': 'str',
        'data_source_instance_id': 'int',
        'disable_dp_alert_host_groups': 'str',
        'data_point_name': 'str',
        'data_point_id': 'int',
        'device_group_id': 'int',
        'parent_device_group_alert_expr_list': 'list[DeviceGroupAlertThresholdInfo]',
        'alerting_disabled_on': 'str',
        'id': 'int',
        'data_source_instance_alias': 'str',
        'device_group_full_path': 'str',
        'alert_expr': 'str',
        'alert_transition_interval': 'int'
    }

    attribute_map = {
        'global_alert_expr': 'globalAlertExpr',
        'alert_clear_interval': 'alertClearInterval',
        'disable_alerting': 'disableAlerting',
        'alert_expr_note': 'alertExprNote',
        'data_point_description': 'dataPointDescription',
        'data_source_instance_id': 'dataSourceInstanceId',
        'disable_dp_alert_host_groups': 'disableDpAlertHostGroups',
        'data_point_name': 'dataPointName',
        'data_point_id': 'dataPointId',
        'device_group_id': 'deviceGroupId',
        'parent_device_group_alert_expr_list': 'parentDeviceGroupAlertExprList',
        'alerting_disabled_on': 'alertingDisabledOn',
        'id': 'id',
        'data_source_instance_alias': 'dataSourceInstanceAlias',
        'device_group_full_path': 'deviceGroupFullPath',
        'alert_expr': 'alertExpr',
        'alert_transition_interval': 'alertTransitionInterval'
    }

    def __init__(self, global_alert_expr=None, alert_clear_interval=None, disable_alerting=None, alert_expr_note=None, data_point_description=None, data_source_instance_id=None, disable_dp_alert_host_groups=None, data_point_name=None, data_point_id=None, device_group_id=None, parent_device_group_alert_expr_list=None, alerting_disabled_on=None, id=None, data_source_instance_alias=None, device_group_full_path=None, alert_expr=None, alert_transition_interval=None):  # noqa: E501
        """DeviceDataSourceInstanceAlertSetting - a model defined in Swagger"""  # noqa: E501

        self._global_alert_expr = None
        self._alert_clear_interval = None
        self._disable_alerting = None
        self._alert_expr_note = None
        self._data_point_description = None
        self._data_source_instance_id = None
        self._disable_dp_alert_host_groups = None
        self._data_point_name = None
        self._data_point_id = None
        self._device_group_id = None
        self._parent_device_group_alert_expr_list = None
        self._alerting_disabled_on = None
        self._id = None
        self._data_source_instance_alias = None
        self._device_group_full_path = None
        self._alert_expr = None
        self._alert_transition_interval = None
        self.discriminator = None

        if global_alert_expr is not None:
            self.global_alert_expr = global_alert_expr
        if alert_clear_interval is not None:
            self.alert_clear_interval = alert_clear_interval
        if disable_alerting is not None:
            self.disable_alerting = disable_alerting
        if alert_expr_note is not None:
            self.alert_expr_note = alert_expr_note
        if data_point_description is not None:
            self.data_point_description = data_point_description
        if data_source_instance_id is not None:
            self.data_source_instance_id = data_source_instance_id
        if disable_dp_alert_host_groups is not None:
            self.disable_dp_alert_host_groups = disable_dp_alert_host_groups
        if data_point_name is not None:
            self.data_point_name = data_point_name
        if data_point_id is not None:
            self.data_point_id = data_point_id
        if device_group_id is not None:
            self.device_group_id = device_group_id
        if parent_device_group_alert_expr_list is not None:
            self.parent_device_group_alert_expr_list = parent_device_group_alert_expr_list
        if alerting_disabled_on is not None:
            self.alerting_disabled_on = alerting_disabled_on
        if id is not None:
            self.id = id
        if data_source_instance_alias is not None:
            self.data_source_instance_alias = data_source_instance_alias
        if device_group_full_path is not None:
            self.device_group_full_path = device_group_full_path
        if alert_expr is not None:
            self.alert_expr = alert_expr
        if alert_transition_interval is not None:
            self.alert_transition_interval = alert_transition_interval

    @property
    def global_alert_expr(self):
        """Gets the global_alert_expr of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The global alert expression for this datapoint  # noqa: E501

        :return: The global_alert_expr of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: str
        """
        return self._global_alert_expr

    @global_alert_expr.setter
    def global_alert_expr(self, global_alert_expr):
        """Sets the global_alert_expr of this DeviceDataSourceInstanceAlertSetting.

        The global alert expression for this datapoint  # noqa: E501

        :param global_alert_expr: The global_alert_expr of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: str
        """

        self._global_alert_expr = global_alert_expr

    @property
    def alert_clear_interval(self):
        """Gets the alert_clear_interval of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The interval of alert clear transition  # noqa: E501

        :return: The alert_clear_interval of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: int
        """
        return self._alert_clear_interval

    @alert_clear_interval.setter
    def alert_clear_interval(self, alert_clear_interval):
        """Sets the alert_clear_interval of this DeviceDataSourceInstanceAlertSetting.

        The interval of alert clear transition  # noqa: E501

        :param alert_clear_interval: The alert_clear_interval of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: int
        """

        self._alert_clear_interval = alert_clear_interval

    @property
    def disable_alerting(self):
        """Gets the disable_alerting of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        Whether or not alerting will be disabled for the datapoint  # noqa: E501

        :return: The disable_alerting of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: bool
        """
        return self._disable_alerting

    @disable_alerting.setter
    def disable_alerting(self, disable_alerting):
        """Sets the disable_alerting of this DeviceDataSourceInstanceAlertSetting.

        Whether or not alerting will be disabled for the datapoint  # noqa: E501

        :param disable_alerting: The disable_alerting of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: bool
        """

        self._disable_alerting = disable_alerting

    @property
    def alert_expr_note(self):
        """Gets the alert_expr_note of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The note associated with the current alert threshold settings  # noqa: E501

        :return: The alert_expr_note of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: str
        """
        return self._alert_expr_note

    @alert_expr_note.setter
    def alert_expr_note(self, alert_expr_note):
        """Sets the alert_expr_note of this DeviceDataSourceInstanceAlertSetting.

        The note associated with the current alert threshold settings  # noqa: E501

        :param alert_expr_note: The alert_expr_note of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: str
        """

        self._alert_expr_note = alert_expr_note

    @property
    def data_point_description(self):
        """Gets the data_point_description of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The description of the datapoint the alert settings apply to  # noqa: E501

        :return: The data_point_description of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: str
        """
        return self._data_point_description

    @data_point_description.setter
    def data_point_description(self, data_point_description):
        """Sets the data_point_description of this DeviceDataSourceInstanceAlertSetting.

        The description of the datapoint the alert settings apply to  # noqa: E501

        :param data_point_description: The data_point_description of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: str
        """

        self._data_point_description = data_point_description

    @property
    def data_source_instance_id(self):
        """Gets the data_source_instance_id of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The id of the DataSource instance alert settings apply to  # noqa: E501

        :return: The data_source_instance_id of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: int
        """
        return self._data_source_instance_id

    @data_source_instance_id.setter
    def data_source_instance_id(self, data_source_instance_id):
        """Sets the data_source_instance_id of this DeviceDataSourceInstanceAlertSetting.

        The id of the DataSource instance alert settings apply to  # noqa: E501

        :param data_source_instance_id: The data_source_instance_id of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: int
        """

        self._data_source_instance_id = data_source_instance_id

    @property
    def disable_dp_alert_host_groups(self):
        """Gets the disable_dp_alert_host_groups of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The group full path lists who disable alert for this datapoint on devicegroup level  # noqa: E501

        :return: The disable_dp_alert_host_groups of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: str
        """
        return self._disable_dp_alert_host_groups

    @disable_dp_alert_host_groups.setter
    def disable_dp_alert_host_groups(self, disable_dp_alert_host_groups):
        """Sets the disable_dp_alert_host_groups of this DeviceDataSourceInstanceAlertSetting.

        The group full path lists who disable alert for this datapoint on devicegroup level  # noqa: E501

        :param disable_dp_alert_host_groups: The disable_dp_alert_host_groups of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: str
        """

        self._disable_dp_alert_host_groups = disable_dp_alert_host_groups

    @property
    def data_point_name(self):
        """Gets the data_point_name of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The name of the datapoint the alert settings apply to  # noqa: E501

        :return: The data_point_name of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: str
        """
        return self._data_point_name

    @data_point_name.setter
    def data_point_name(self, data_point_name):
        """Sets the data_point_name of this DeviceDataSourceInstanceAlertSetting.

        The name of the datapoint the alert settings apply to  # noqa: E501

        :param data_point_name: The data_point_name of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: str
        """

        self._data_point_name = data_point_name

    @property
    def data_point_id(self):
        """Gets the data_point_id of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The id of the Datapoint alert settings apply to  # noqa: E501

        :return: The data_point_id of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: int
        """
        return self._data_point_id

    @data_point_id.setter
    def data_point_id(self, data_point_id):
        """Sets the data_point_id of this DeviceDataSourceInstanceAlertSetting.

        The id of the Datapoint alert settings apply to  # noqa: E501

        :param data_point_id: The data_point_id of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: int
        """

        self._data_point_id = data_point_id

    @property
    def device_group_id(self):
        """Gets the device_group_id of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The ID of the device group  # noqa: E501

        :return: The device_group_id of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: int
        """
        return self._device_group_id

    @device_group_id.setter
    def device_group_id(self, device_group_id):
        """Sets the device_group_id of this DeviceDataSourceInstanceAlertSetting.

        The ID of the device group  # noqa: E501

        :param device_group_id: The device_group_id of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: int
        """

        self._device_group_id = device_group_id

    @property
    def parent_device_group_alert_expr_list(self):
        """Gets the parent_device_group_alert_expr_list of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        Device group alert expression list base on the priority. The first is the highest priority and effected on this instance  # noqa: E501

        :return: The parent_device_group_alert_expr_list of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: list[DeviceGroupAlertThresholdInfo]
        """
        return self._parent_device_group_alert_expr_list

    @parent_device_group_alert_expr_list.setter
    def parent_device_group_alert_expr_list(self, parent_device_group_alert_expr_list):
        """Sets the parent_device_group_alert_expr_list of this DeviceDataSourceInstanceAlertSetting.

        Device group alert expression list base on the priority. The first is the highest priority and effected on this instance  # noqa: E501

        :param parent_device_group_alert_expr_list: The parent_device_group_alert_expr_list of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: list[DeviceGroupAlertThresholdInfo]
        """

        self._parent_device_group_alert_expr_list = parent_device_group_alert_expr_list

    @property
    def alerting_disabled_on(self):
        """Gets the alerting_disabled_on of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The datapoint is effected alert disabled by which group  # noqa: E501

        :return: The alerting_disabled_on of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: str
        """
        return self._alerting_disabled_on

    @alerting_disabled_on.setter
    def alerting_disabled_on(self, alerting_disabled_on):
        """Sets the alerting_disabled_on of this DeviceDataSourceInstanceAlertSetting.

        The datapoint is effected alert disabled by which group  # noqa: E501

        :param alerting_disabled_on: The alerting_disabled_on of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: str
        """

        self._alerting_disabled_on = alerting_disabled_on

    @property
    def id(self):
        """Gets the id of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The id of this alert setting  # noqa: E501

        :return: The id of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: int
        """
        return self._id

    @id.setter
    def id(self, id):
        """Sets the id of this DeviceDataSourceInstanceAlertSetting.

        The id of this alert setting  # noqa: E501

        :param id: The id of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: int
        """

        self._id = id

    @property
    def data_source_instance_alias(self):
        """Gets the data_source_instance_alias of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The alias (name) of the DataSource instance the alert settings apply to  # noqa: E501

        :return: The data_source_instance_alias of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: str
        """
        return self._data_source_instance_alias

    @data_source_instance_alias.setter
    def data_source_instance_alias(self, data_source_instance_alias):
        """Sets the data_source_instance_alias of this DeviceDataSourceInstanceAlertSetting.

        The alias (name) of the DataSource instance the alert settings apply to  # noqa: E501

        :param data_source_instance_alias: The data_source_instance_alias of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: str
        """

        self._data_source_instance_alias = data_source_instance_alias

    @property
    def device_group_full_path(self):
        """Gets the device_group_full_path of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The full path of the device group  # noqa: E501

        :return: The device_group_full_path of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: str
        """
        return self._device_group_full_path

    @device_group_full_path.setter
    def device_group_full_path(self, device_group_full_path):
        """Sets the device_group_full_path of this DeviceDataSourceInstanceAlertSetting.

        The full path of the device group  # noqa: E501

        :param device_group_full_path: The device_group_full_path of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: str
        """

        self._device_group_full_path = device_group_full_path

    @property
    def alert_expr(self):
        """Gets the alert_expr of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The thresholds that should be associated with the datapoint. Note that you need to have a space between the operator and each threshold (e.g. > 1 2 3)  # noqa: E501

        :return: The alert_expr of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: str
        """
        return self._alert_expr

    @alert_expr.setter
    def alert_expr(self, alert_expr):
        """Sets the alert_expr of this DeviceDataSourceInstanceAlertSetting.

        The thresholds that should be associated with the datapoint. Note that you need to have a space between the operator and each threshold (e.g. > 1 2 3)  # noqa: E501

        :param alert_expr: The alert_expr of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: str
        """

        self._alert_expr = alert_expr

    @property
    def alert_transition_interval(self):
        """Gets the alert_transition_interval of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501

        The interval of alert transition  # noqa: E501

        :return: The alert_transition_interval of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :rtype: int
        """
        return self._alert_transition_interval

    @alert_transition_interval.setter
    def alert_transition_interval(self, alert_transition_interval):
        """Sets the alert_transition_interval of this DeviceDataSourceInstanceAlertSetting.

        The interval of alert transition  # noqa: E501

        :param alert_transition_interval: The alert_transition_interval of this DeviceDataSourceInstanceAlertSetting.  # noqa: E501
        :type: int
        """

        self._alert_transition_interval = alert_transition_interval

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(DeviceDataSourceInstanceAlertSetting, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, DeviceDataSourceInstanceAlertSetting):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other