import httplib2
import os
import time

from apiclient import discovery
from oauth2client import client
from oauth2client import tools
from oauth2client.service_account import ServiceAccountCredentials
from oauth2client.file import Storage
import Common

SCOPES = ['https://www.googleapis.com/auth/spreadsheets',
          'https://www.googleapis.com/auth/drive',
          'https://www.googleapis.com/auth/drive.appdata',
          'https://www.googleapis.com/auth/drive.metadata',
          'https://www.googleapis.com/auth/drive.scripts',
          'https://www.googleapis.com/auth/drive.file']
CLIENT_SECRET_FILE = 'client_secret.json'
SERVICE_ACCOUNT_FILE = 'service.json'
DISCOVERY_URL_SHEETS = 'https://sheets.googleapis.com/$discovery/rest?version=v4'

MASTER_DIRECTORY_ID = ""
APP_PROPERTY_KEY = "master"
APP_PROPERTIES = {APP_PROPERTY_KEY: "1"}

class GoogleDriveHelper:
    _credentials = None
    _drive = None
    _sheets  = None
    _target_files = []

    def __init__(self):
        self._get_drive_client()

    def _get_credentials(self):
        if self._credentials == None:
            self._credentials = ServiceAccountCredentials.from_json_keyfile_name(
                SERVICE_ACCOUNT_FILE,
                scopes=SCOPES)
        return self._credentials

    def _get_drive_client(self):
        if self._drive == None:
            http = self._get_credentials().authorize(httplib2.Http())
            self._drive = discovery.build('drive', 'v3', http=http)
        return self._drive

    def _get_sheets_client(self):
        if self._sheets == None:
            http = self._get_credentials().authorize(httplib2.Http())
            self._sheets = discovery.build('sheets', 'v4', http=http, discoveryServiceUrl=DISCOVERY_URL_SHEETS)
        return self._sheets

    def get_master_file(self, include=[], exclude=['ScenarioMaster']):
        query = ("appProperties has { key='"+ APP_PROPERTY_KEY +"' and value='1' }"
            " and trashed = false ")

        # exclude ファイル反映
        for master in exclude:
            query = query + " and not name contains '" + master + "'"

        # include ファイル反映
        tmp = None
        for master in include:
            if tmp == None:
                tmp = "name contains '" + master + "'"
            else:
                tmp += " or name contains '" + master + "'"

        if tmp != None:
            query = query + " and (" + tmp + ")"

        drive = self._get_drive_client()
        result = drive.files().list(includeTeamDriveItems=True,
                                    supportsTeamDrives=True,
                                    fields='nextPageToken, files(id, name, mimeType)',
                                    pageSize=1000,
                                    q=query).execute()
        files = result.get('files',[])
        return files

    def get_asuka_general_master_file(self):
        return self.get_master_file(exclude=['ScenarioMaster'])

    def get_asuka_news_master_file(self):
        return self.get_master_file(include=['NewsMaster'])

    def get_file_list_by_id(self, path_id):
        """ Google Drive からファイル一覧を取得します

        Args:
            path_id: string     ディレクトリの id

        Returns:
            files 配列
        """
        drive = self._get_drive_client()
        result = drive.files().list(includeTeamDriveItems=True,
                                    supportsTeamDrives=True,
                                    fields='nextPageToken, files(id, name, mimeType, appProperties)',
                                    pageSize=1000,
                                    q="'" + path_id + "' in parents and trashed = false ").execute()
        #print(result)
        files = result.get('files',[])
        return files

    def get_files_recursive(self, path_id, file_list):
        files = self.get_file_list_by_id(path_id)
        for file in files:
            if file['mimeType'] == 'application/vnd.google-apps.folder':
                #if file['name'] != 'Scenarios':
                self.get_files_recursive(file['id'], file_list)
            else:
                file_list.append(file)

    def add_app_properties(self, file_id):
        drive = self._get_drive_client
        result = drive.files().update(fileId=file_id,
                                supportsTeamDrives=True,
                                fields='id, name, appProperties',
                                body={"appProperties": APP_PROPERTIES}).execute()
        print(result)


    def add_app_property(self, file_id):
        self._target_files = []
        self.get_files_recursive(file_id, self._target_files)

        for file in self._target_files:
            if 'appProperties' not in file:
                self.add_app_properties(file['id'])
            elif 'master' not in file['appProperties']:
                self.add_app_properties(file['id'])

    def get_values_by_id(self, spreadsheet_id):
        sheets = self._get_sheets_client()
        result = sheets.spreadsheets().values().get(
            spreadsheetId=spreadsheet_id, range="Sheet1").execute()
        values = result.get('values', [])
        return values

    def get_spreadsheet(self, spreadsheet_id, include_data = False):
        sheets = self._get_sheets_client()
        request = sheets.spreadsheets().get(spreadsheetId=spreadsheet_id,includeGridData = include_data)
        response = request.execute()
        return response

def main():
    instance = GoogleDriveHelper()

    # ディレクトリ配下のファイル一覧取得
    #file_list = []
    #print(time.ctime())
    #instance.get_files_recursive(MASTER_DIRECTORY_ID, file_list)
    #print(file_list)
    #print(time.ctime())

    # ディレクトリ配下のファイルの appProperties に、 "master":"1" を設定する
    # Master ディレクトリを指定して、Master ファイル全てに "master":"1" を設定する
    #print(time.ctime())
    #instance.add_app_property(MASTER_DIRECTORY_ID)
    #print(time.ctime())

    # "master":"1" が設定されているファイル一覧を取得する
    #print(time.ctime())
    #self._target_file = self.get_master_file()
    #print(time.ctime())

    instance.get_asuka_news_master_file()


if __name__ == '__main__':
    main()