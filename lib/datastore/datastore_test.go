package datastore_test

import (
	"errors"
	"os"

	"code.cloudfoundry.org/silk/lib/datastore"

	libfakes "code.cloudfoundry.org/silk/lib/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Datastore", func() {
	var (
		handle   string
		ip       string
		store    *datastore.Store
		metadata map[string]interface{}

		serializer *libfakes.Serializer
		locker     *libfakes.FileLocker
		lockedFile *os.File
		filePath   string
	)

	BeforeEach(func() {
		handle = "some-handle"
		ip = "192.168.0.100"
		filePath = "file"
		locker = &libfakes.FileLocker{}
		serializer = &libfakes.Serializer{}
		metadata = map[string]interface{}{
			"AppID":         "some-appid",
			"OrgID":         "some-orgid",
			"PolicyGroupID": "some-policygroupid",
			"SpaceID":       "some-spaceid",
			"randomKey":     "randomValue",
		}

		store = &datastore.Store{
			Serializer: serializer,
			Locker:     locker,
		}

		lockedFile = &os.File{}
		locker.OpenReturns(lockedFile, nil)
	})

	Context("when adding an entry to store", func() {
		It("deserializes the data from the file", func() {
			err := store.Add(filePath, handle, ip, metadata)
			Expect(err).NotTo(HaveOccurred())

			Expect(locker.OpenCallCount()).To(Equal(1))
			Expect(serializer.DecodeAllCallCount()).To(Equal(1))
			Expect(serializer.EncodeAndOverwriteCallCount()).To(Equal(1))

			Expect(locker.OpenArgsForCall(0)).To(Equal(filePath))

			file, _ := serializer.DecodeAllArgsForCall(0)
			Expect(file).To(Equal(lockedFile))

			_, actual := serializer.EncodeAndOverwriteArgsForCall(0)
			expected := map[string]datastore.Container{
				handle: datastore.Container{
					Handle:   handle,
					IP:       ip,
					Metadata: metadata,
				},
			}
			Expect(actual).To(Equal(expected))
		})

		Context("when handle is not valid", func() {
			It("wraps and returns the error", func() {
				err := store.Add(filePath, "", ip, metadata)
				Expect(err).To(MatchError("invalid handle"))
			})
		})

		Context("when input IP is not valid", func() {
			It("wraps and returns the error", func() {
				err := store.Add(filePath, handle, "invalid-ip", metadata)
				Expect(err).To(MatchError("invalid ip: invalid-ip"))
			})
		})

		Context("when file locker fails to open", func() {
			BeforeEach(func() {
				locker.OpenReturns(nil, errors.New("potato"))
			})
			It("wraps and returns the error", func() {
				err := store.Add(filePath, handle, ip, metadata)
				Expect(err).To(MatchError("open lock: potato"))
			})
		})

		Context("when serializer fails to decode", func() {
			BeforeEach(func() {
				serializer.DecodeAllReturns(errors.New("potato"))
			})
			It("wraps and returns the error", func() {
				err := store.Add(filePath, handle, ip, metadata)
				Expect(err).To(MatchError("decoding file: potato"))
			})
		})

		Context("when serializer fails to encode", func() {
			BeforeEach(func() {
				serializer.EncodeAndOverwriteReturns(errors.New("potato"))
			})
			It("wraps and returns the error", func() {
				err := store.Add(filePath, handle, ip, metadata)
				Expect(err).To(MatchError("encode and overwrite: potato"))
			})
		})

	})

	Context("when deleting an entry from store", func() {

		It("deserializes the data from the file", func() {
			_, err := store.Delete(filePath, handle)
			Expect(err).NotTo(HaveOccurred())

			Expect(locker.OpenCallCount()).To(Equal(1))
			Expect(serializer.DecodeAllCallCount()).To(Equal(1))
			Expect(serializer.EncodeAndOverwriteCallCount()).To(Equal(1))

			Expect(locker.OpenArgsForCall(0)).To(Equal(filePath))

			file, _ := serializer.DecodeAllArgsForCall(0)
			Expect(file).To(Equal(lockedFile))

			_, actual := serializer.EncodeAndOverwriteArgsForCall(0)
			Expect(actual).ToNot(HaveKey(handle))
		})

		Context("when handle is not valid", func() {
			It("wraps and returns the error", func() {
				_, err := store.Delete(filePath, "")
				Expect(err).To(MatchError("invalid handle"))
			})
		})

		Context("when file locker fails to open", func() {
			BeforeEach(func() {
				locker.OpenReturns(nil, errors.New("potato"))
			})
			It("wraps and returns the error", func() {
				_, err := store.Delete(filePath, handle)
				Expect(err).To(MatchError("open lock: potato"))
			})
		})

		Context("when serializer fails to decode", func() {
			BeforeEach(func() {
				serializer.DecodeAllReturns(errors.New("potato"))
			})
			It("wraps and returns the error", func() {
				_, err := store.Delete(filePath, handle)
				Expect(err).To(MatchError("decoding file: potato"))
			})
		})

		Context("when serializer fails to encode", func() {
			BeforeEach(func() {
				serializer.EncodeAndOverwriteReturns(errors.New("potato"))
			})
			It("wraps and returns the error", func() {
				_, err := store.Delete(filePath, handle)
				Expect(err).To(MatchError("encode and overwrite: potato"))
			})
		})

	})

	Context("when reading from datastore", func() {
		It("deserializes the data from the file", func() {
			data, err := store.ReadAll(filePath)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).NotTo(BeNil())

			Expect(locker.OpenCallCount()).To(Equal(1))
			Expect(serializer.DecodeAllCallCount()).To(Equal(1))
			Expect(serializer.EncodeAndOverwriteCallCount()).To(Equal(0))

			Expect(locker.OpenArgsForCall(0)).To(Equal(filePath))

			file, _ := serializer.DecodeAllArgsForCall(0)
			Expect(file).To(Equal(lockedFile))
		})

		Context("when file locker fails to open", func() {
			BeforeEach(func() {
				locker.OpenReturns(nil, errors.New("potato"))
			})
			It("wraps and returns the error", func() {
				_, err := store.ReadAll(filePath)
				Expect(err).To(MatchError("open lock: potato"))
			})
		})

		Context("when serializer fails to decode", func() {
			BeforeEach(func() {
				serializer.DecodeAllReturns(errors.New("potato"))
			})
			It("wraps and returns the error", func() {
				_, err := store.ReadAll(filePath)
				Expect(err).To(MatchError("decoding file: potato"))
			})
		})
	})
})
