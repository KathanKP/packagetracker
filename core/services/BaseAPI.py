"""
This is the base api class which contains basic functionality about an API call for a service.
Each object deals with one tracking number/ one package
It will be the parent class for individual service level classes
"""


class BaseAPI:

    def __init__(self, tracking_number):
        self.tracking_number = tracking_number

    def get_tracking_number(self):
        return self.tracking_number
